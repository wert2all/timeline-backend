package auth

import (
	"timeline/backend/db/model/user"

	"google.golang.org/api/idtoken"
)

func From(payload idtoken.Payload) user.GoogleUser {
	return user.GoogleUser{
		UUID:   payload.Claims["sub"].(string),
		Name:   payload.Claims["name"].(string),
		Email:  payload.Claims["email"].(string),
		Avatar: payload.Claims["picture"].(string),
	}
}
