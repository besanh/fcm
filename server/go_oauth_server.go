package server

import (
	"fcm/common/env"
	"fcm/common/log"
	"fcm/common/response"
	"net/http"

	"github.com/gin-gonic/gin"
	oauthErrror "github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	oauthServer "github.com/go-oauth2/oauth2/v4/server"
	store "github.com/go-oauth2/oauth2/v4/store"
	oredis "github.com/go-oauth2/redis/v4"
	redisv8 "github.com/go-redis/redis/v8"
)

func (server *Server) NewOAuthServer() {
	manager := manage.NewDefaultManager()

	// Modify token, TTL
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000",
		Secret: "000",
		Domain: "http://localhost:8000",
	})

	manager.MapClientStorage(clientStore)
	manager.MapTokenStorage(oredis.NewRedisStore(&redisv8.Options{
		Addr: env.GetStringENV("REDIS_HOST", "localhost:6379"),
		DB:   env.GetIntENV("REDIS_DB", 0),
	}))

	srv := oauthServer.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(oauthServer.ClientFormHandler)
	srv.UserAuthorizationHandler = func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
		return "0000", nil
	}

	srv.SetInternalErrorHandler(func(err error) (re *oauthErrror.Response) {
		log.Error("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *oauthErrror.Response) {
		log.Error("Response Error:", re.Error.Error())
	})
	server.Engine.GET("/authorize", func(c *gin.Context) {
		err := srv.HandleAuthorizeRequest(c.Writer, c.Request)
		if err != nil {
			c.JSON(response.NotFoundMsg(err.Error()))
			return
		}
	})
	server.Engine.POST("/token", func(c *gin.Context) {
		err := srv.HandleTokenRequest(c.Writer, c.Request)
		if err != nil {
			c.JSON(response.NotFoundMsg(err.Error()))
			return
		}
	})

}
