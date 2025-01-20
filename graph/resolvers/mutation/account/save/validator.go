package save

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
	ValidSaveAccountArguments struct {
		account  *ent.Account
		name     string
		avatarID *int
	}
	validatorImpl struct {
		userModel     user.UserModel
		userExtractor domainUser.UserExtractor
	}
)

// Validate implements resolvers.Validator.
func (v validatorImpl) Validate(ctx context.Context, arguments resolvers.Arguments[SaveAccountArguments]) (resolvers.ValidArguments[ValidSaveAccountArguments], error) {
	token := appContext.GetToken(ctx)
	user, errExtraction := v.userExtractor.ExtractUserFromToken(ctx, token)
	if errExtraction != nil {
		return nil, errors.New("could save account: " + errExtraction.Error())
	}
	args := arguments.GetArguments()
	account, err := v.userModel.GetUserAccount(args.accountID, user.ID)
	if err != nil {
		return nil, errors.New("could save account: " + err.Error())
	}
	return ValidSaveAccountArguments{
		account:  account,
		name:     args.name,
		avatarID: args.avatarID,
	}, nil
}

// GetArguments implements resolvers.ValidArguments.
func (v ValidSaveAccountArguments) GetArguments() ValidSaveAccountArguments { return v }

func NewValidator(userModel user.UserModel, userExtractor domainUser.UserExtractor) resolvers.Validator[SaveAccountArguments, ValidSaveAccountArguments] {
	return validatorImpl{
		userModel:     userModel,
		userExtractor: userExtractor,
	}
}
