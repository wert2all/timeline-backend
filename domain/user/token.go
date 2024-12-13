package user

import (
	"context"
	"errors"

	"timeline/backend/db/model/user"

	"google.golang.org/api/idtoken"
)

type UserExtractor struct {
	googleClientID string
	userModel      user.Authorize
}

func (u UserExtractor) ExtractUserFromToken(ctx context.Context, token *string) (*user.CheckOrCreate, error) {
	if token == nil {
		return nil, errors.New("token is nil")
	}
	payload, err := idtoken.Validate(ctx, *token, u.googleClientID)
	if err != nil {
		return nil, errors.New("Invalid token")
	}

	someUser := user.NewSomeUser(
		payload.Claims["sub"].(string),
		payload.Claims["name"].(string),
		payload.Claims["email"].(string),
		payload.Claims["picture"].(string),
	)

	userCheck, err := u.userModel.CheckOrCreate(ctx, someUser)
	if err != nil {
		return nil, errors.New("Blocked")
	}
	return userCheck, nil
}

func NewUserExtractor(googleClientID string, userModel user.Authorize) UserExtractor {
	return UserExtractor{
		googleClientID: googleClientID,
		userModel:      userModel,
	}
}
