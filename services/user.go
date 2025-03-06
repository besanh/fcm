package services

import (
	"context"
	"fcm/common/cache"
	"fcm/common/constant"
	"fcm/common/log"
	"fcm/common/util"
	"fcm/models"
	"fcm/pkgs/oauth"
	"fcm/repositories"
	"fmt"
	"net/url"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"
)

type (
	IUser interface {
		Login(ctx context.Context) (url string)
		OAuth2Callback(ctx context.Context, callbackData *models.OAuth2Callback) (token string, err error)
	}

	User struct {
		userRepo     repositories.IUser
		oauth2Client oauth.IOAuth2
	}
)

func NewUser(userRepo repositories.IUser, oauth2Client oauth.IOAuth2) IUser {
	return &User{
		userRepo:     userRepo,
		oauth2Client: oauth2Client,
	}
}

/*
 * Login and return url(google app)
 */
func (s *User) Login(ctx context.Context) (callbackUrl string) {
	// use PKCE to protect against CSRF attacks
	// https://www.ietf.org/archive/id/draft-ietf-oauth-security-topics-22.html#name-countermeasures-6
	verifier := oauth2.GenerateVerifier()
	// Generate PKCE values.
	challenge := util.GenerateCodeChallenge(verifier)

	// Generate the authorization URL with PKCE parameters.
	authURL, err := url.Parse(s.oauth2Client.AuthCodeUrl(OAUTH2_STATE, verifier))
	if err != nil {
		log.Error(err)
		return
	}

	q := authURL.Query()
	q.Set("code_challenge", challenge)
	q.Set("code_challenge_method", "S256")
	authURL.RawQuery = q.Encode()

	// Caching
	redisKey := fmt.Sprintf("pkce:%s", OAUTH2_STATE)
	if err := cache.RCache.Set(redisKey, verifier, 3*time.Minute); err != nil {
		log.Error(err)
		return
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	callbackUrl = s.oauth2Client.AuthCodeUrl(OAUTH2_STATE, verifier)
	return
}

func (s *User) OAuth2Callback(ctx context.Context, callbackData *models.OAuth2Callback) (token string, err error) {
	// PKCE
	verifierCache := cache.RCache.Get(fmt.Sprintf("pkce:%s", callbackData.State))
	if verifierCache == nil {
		err = fmt.Errorf("invalid state: %s", callbackData.State)
		return
	}
	verifier := verifierCache.(string)

	// Exchange the code for a token
	userInfo, err := s.oauth2Client.Exchange(ctx, callbackData.Code, oauth2.SetAuthURLParam("code_verifier", verifier))
	if err != nil {
		log.Error(err)
		return
	}
	// Calculate TTL for the access token based on its expiry time
	ttl := time.Until(userInfo.Expiry)

	refreshToken, err := util.Encrypt(userInfo.RefreshToken)
	if err != nil {
		log.Error(err)
		return
	}

	user := &models.User{
		GBase:                 models.InitBase(),
		Status:                constant.USER_STATUS_ACTIVE,
		RefreshTokenEncrypted: refreshToken,
	}

	// Start transaction redis
	redisTx := cache.RCache.TxPineLine()
	cache.RCache.TxSet(ctx, redisTx, OAUTH2_TOKEN, userInfo.AccessToken, ttl)

	// Start transaction mongodb
	mongoSession, err := s.userRepo.StartSession()
	if err != nil {
		log.Error(err)
		return
	}

	defer mongoSession.EndSession(ctx)

	// Transaction mongodb
	if err = mongo.WithSession(ctx, mongoSession, func(sc mongo.SessionContext) (err error) {
		if err = s.userRepo.StartTransaction(mongoSession); err != nil {
			return
		}

		if err = s.userRepo.Insert(ctx, user); err != nil {
			s.userRepo.AbortTransaction(ctx, mongoSession)
			return err
		}

		if err = s.userRepo.CommitTransaction(ctx, mongoSession); err != nil {
			s.userRepo.AbortTransaction(ctx, mongoSession)
			return
		}

		return
	}); err != nil {
		log.Error(err)
		return
	}

	// Commit redis
	if _, err = redisTx.Exec(ctx); err != nil {
		log.Error(err)
		return
	}

	token = userInfo.AccessToken

	return
}
