package repositories

import (
	"context"
	"fcm/common/log"
	"fcm/models"
	"fcm/pkgs/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	IUser interface {
		IRepoGeneric[models.User]
	}
	User struct {
		RepoGeneric[models.User]
	}
)

var UserRepo IUser

/*
 * Declare new repo with collection(table)
 */
func NewUser(db *mongodb.IMongoDBClient) IUser {
	repo := &User{
		RepoGeneric: RepoGeneric[models.User]{
			DB:         *db,
			Collection: "users",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "name", Value: repo.Collection}}
	collections, err := repo.DB.DB().ListCollectionNames(ctx, filter)
	if err != nil {
		log.Errorf("failed to list collections: %v", err)
		panic(err)
	}
	if len(collections) == 0 {
		err = repo.DB.DB().CreateCollection(ctx, repo.Collection)
		if err != nil {
			log.Errorf("failed to create collection: %v", err)
		} else {
			log.Debugf("collection '%s' created", repo.Collection)
		}
	}

	return repo
}
