package getevent

import (
	"context"

	appContext "timeline/backend/app/context"
	"timeline/backend/db/model/user"
	domainUser "timeline/backend/domain/user"
	"timeline/backend/ent"
	"timeline/backend/graph/resolvers"
)

type (
	validatorImpl struct {
		userModel     user.UserModel
		userExtractor domainUser.UserExtractor
	}
	ValidGetEventArguments struct {
		eventID int
		account *ent.Account
	}
)

// GetArguments implements resolvers.ValidArguments.
func (v ValidGetEventArguments) GetArguments() ValidGetEventArguments { return v }

// Validate implements resolvers.Validator.
func (v validatorImpl) Validate(context context.Context, arg resolvers.Arguments[GetEventArguments]) (resolvers.ValidArguments[ValidGetEventArguments], error) {
	return ValidGetEventArguments{eventID: arg.GetArguments().eventID, account: v.extractAccount(context, arg.GetArguments().accountID)}, nil
}

func (v validatorImpl) extractAccount(ctx context.Context, requestAccountID *int) *ent.Account {
	token := appContext.GetToken(ctx)
	user, errExtraction := v.userExtractor.ExtractUserFromToken(ctx, token)
	if errExtraction != nil {
		return nil
	}
	if requestAccountID != nil {
		account, err := v.userModel.GetUserAccount(*requestAccountID, user.ID)
		if err != nil {
			return nil
		}
		return account
	}
	return nil
}

func NewValidator(userModel user.UserModel, userExtractor domainUser.UserExtractor) resolvers.Validator[GetEventArguments, ValidGetEventArguments] {
	return validatorImpl{userExtractor: userExtractor, userModel: userModel}
}
