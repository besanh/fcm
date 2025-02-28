package services

import (
	"context"
	"fcm/models"
	"fcm/repositories"
)

type (
	IUser interface {
		InsertUser(ctx context.Context, request models.UserPostRequest) (err error)
	}

	User struct {
		userRepo repositories.IUser
	}
)

func NewUser(userRepo repositories.IUser) IUser {
	return &User{
		userRepo: userRepo,
	}
}

func (s *User) InsertUser(ctx context.Context, request models.UserPostRequest) (err error) {
	return
}
