package add

import (
	"context"
	"errors"

	appContext "timeline/backend/app/context"
	"timeline/backend/db/model/user"
	domainUser "timeline/backend/domain/user"
	"timeline/backend/ent"
	"timeline/backend/graph/resolvers"
)

type (
	ValidAddAccountArguments struct {
		user *ent.User
		name string
	}
	validatorImpl struct {
		userExtractor domainUser.UserExtractor
		userModel     user.UserModel
	}
)

// GetArguments implements resolvers.ValidArguments.
func (v ValidAddAccountArguments) GetArguments() ValidAddAccountArguments { return v }

// Validate implements resolvers.Validator.
func (v validatorImpl) Validate(ctx context.Context, arguments resolvers.Arguments[AddAccountArguments]) (resolvers.ValidArguments[ValidAddAccountArguments], error) {
	token := appContext.GetToken(ctx)
	extractedUser, errExtraction := v.userExtractor.ExtractUserFromToken(ctx, token)
	if errExtraction != nil {
		return nil, errors.New("could add account: " + errExtraction.Error())
	}
	user, err := v.userModel.GetUser(extractedUser.ID)
	if err != nil {
		return nil, errors.New("could add account: " + err.Error())
	}
	return ValidAddAccountArguments{name: arguments.GetArguments().Name, user: user}, nil
}

func NewValidator(userExtractor domainUser.UserExtractor, userModel user.UserModel) resolvers.Validator[AddAccountArguments, ValidAddAccountArguments] {
	return validatorImpl{userExtractor: userExtractor, userModel: userModel}
}
