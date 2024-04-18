package auth

import (
	"google.golang.org/api/idtoken"
)

type GoogleUser struct {
	uuid   string
	name   string
	email  string
	avatar string
}

func From(payload idtoken.Payload) GoogleUser {
	return GoogleUser{
		uuid:   payload.Claims["sub"].(string),
		name:   payload.Claims["name"].(string),
		email:  payload.Claims["email"].(string),
		avatar: payload.Claims["picture"].(string),
	}
}
