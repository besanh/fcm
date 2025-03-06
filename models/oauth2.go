package models

type (
	OAuth2Callback struct {
		Code  string `json:"code"`
		State string `json:"state"`
		Scope string `json:"scope"`
	}
)
