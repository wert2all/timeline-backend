package user

import (
	"context"

	"timeline/backend/db/model/user"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/api/idtoken"
)

type UserExtractor struct {
	googleClientID string
	userModel      user.Authorize
}

func (u UserExtractor) ExtractUserFromToken(ctx context.Context, token *string) (*user.CheckOrCreate, error) {
	if token == nil {
		return nil, &gqlerror.Error{
			Message: "empty token",
			Path:    graphql.GetPath(ctx),
			Extensions: map[string]interface{}{
				"code": "auth/empty_token",
			},
		}
	}
	payload, err := idtoken.Validate(ctx, *token, u.googleClientID)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: "invalid token",
			Path:    graphql.GetPath(ctx),
			Extensions: map[string]interface{}{
				"code": "auth/invalid_token",
			},
		}
	}

	someUser := user.NewSomeUser(
		payload.Claims["sub"].(string),
		payload.Claims["name"].(string),
		payload.Claims["email"].(string),
		payload.Claims["picture"].(string),
	)

	userCheck, err := u.userModel.CheckOrCreate(ctx, someUser)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: "blocked",
			Path:    graphql.GetPath(ctx),
			Extensions: map[string]interface{}{
				"code": "auth/blocked_user",
			},
		}
	}
	return userCheck, nil
}

func NewUserExtractor(googleClientID string, userModel user.Authorize) UserExtractor {
	return UserExtractor{
		googleClientID: googleClientID,
		userModel:      userModel,
	}
}
