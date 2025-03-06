package models

type (
	User struct {
		*GBase
		RefreshTokenEncrypted string   `json:"refresh_token_encrypted" bson:"refresh_token_encrypted"`
		Status                string   `json:"status" bson:"status"`
		Scope                 []string `json:"scope" bson:"scope"`
	}
)
