package models

import (
	"time"
)

type (
	GModel interface {
		GetId() string
		SetId(id string)
		SetCreatedAt(t time.Time)
		SetUpdatedAt(t time.Time)
	}

	GBase struct {
		Id        string    `json:"id" bson:"_id"`
		CreatedAt time.Time `json:"created_at" bson:"created_at"`
		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	}
)

func (b *GBase) GetId() string {
	return b.Id
}

func (b *GBase) SetId(id string) {
	b.Id = id
}

func (b *GBase) SetCreatedAt(t time.Time) {
	b.CreatedAt = t
}

func (b *GBase) SetUpdatedAt(t time.Time) {
	b.UpdatedAt = t
}
