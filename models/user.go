package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	User struct {
		*GBase
		ClientId primitive.ObjectID `json:"client_id" bson:"client_id"`
		Secret   string             `json:"secret" bson:"secret"`
		Status   string             `json:"status" bson:"status"`
		Scope    []string           `json:"scope" bson:"scope"`
	}

	UserPostRequest struct {
		Id       primitive.ObjectID `json:"id" bson:"id"`
		ClientId primitive.ObjectID `json:"client_id" bson:"client_id"`
		Secret   string             `json:"secret" bson:"secret"`
		Status   string             `json:"status" bson:"status"`
		Scope    []string           `json:"scope" bson:"scope"`
	}
)
