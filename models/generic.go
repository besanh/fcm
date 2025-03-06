package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	GModel interface {
		GetId() string
		SetId(id string)
		SetCreatedAt(t time.Time)
		SetUpdatedAt(t time.Time)
	}

	GBase struct {
		Id        primitive.ObjectID `json:"id" bson:"_id"`
		CreatedAt time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	}
)

func (b *GBase) GetId() string {
	return b.Id.Hex()
}

func (b *GBase) SetId(id string) {
	b.Id, _ = primitive.ObjectIDFromHex(id)
}

func (b *GBase) SetCreatedAt(t time.Time) {
	b.CreatedAt = t
}

func (b *GBase) SetUpdatedAt(t time.Time) {
	b.UpdatedAt = t
}

func InitBase() *GBase {
	return &GBase{
		Id:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
