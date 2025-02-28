package repositories

import (
	"fcm/models"
	"fcm/pkgs/mongodb"
)

type (
	IUser interface {
		IRepoGeneric[models.User]
	}
	User struct {
		RepoGeneric[models.User]
	}
)

func NewUser(db mongodb.IMongoDBClient) IUser {
	return &User{
		RepoGeneric: RepoGeneric[models.User]{
			DB:         db,
			Collection: "users",
		},
	}
}
