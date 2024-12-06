package settings

import (
	"context"
	"errors"

	"timeline/backend/db/model/user"
	"timeline/backend/graph/model"
	"timeline/backend/graph/resolvers"
)

type (
	SaveSettingsArguments struct {
		accountID int
		userID    int
		settings  []*model.AccountSettingInput
	}
	ValidSaveSettingsArguments struct {
		accountID int
		settings  map[string]string
	}
	SaveSettingsArgumentFactory struct{}

	validatorImpl struct {
		userModel user.UserModel
	}
	resolverImpl struct{}
)

// GetArguments implements resolvers.ValidArguments.
func (v ValidSaveSettingsArguments) GetArguments() ValidSaveSettingsArguments { return v }

// GetArguments implements resolvers.Arguments.
func (s SaveSettingsArguments) GetArguments() SaveSettingsArguments { return s }

// Resolve implements resolvers.Resolver.
func (r resolverImpl) Resolve(ctx context.Context, arguments resolvers.ValidArguments[ValidSaveSettingsArguments]) (model.Status, error) {
	panic("unimplemented")
}

// Validate implements resolvers.Validator.
func (v validatorImpl) Validate(ctx context.Context, arguments resolvers.Arguments[SaveSettingsArguments]) (resolvers.ValidArguments[ValidSaveSettingsArguments], error) {
	args := arguments.GetArguments()
	_, err := v.userModel.GetUserAccount(args.accountID, args.userID)
	if err != nil {
		return nil, errors.New("could not save settings")
	}

	if len(args.settings) == 0 {
		return nil, errors.New("could not save empty settings")
	}
	settings := map[string]string{}
	for _, s := range args.settings {
		if s.Value != nil {
			settings[s.Key] = *s.Value
		} else {
			settings[s.Key] = ""
		}
	}
	return ValidSaveSettingsArguments{
		accountID: args.accountID,
		settings:  settings,
	}, nil
}

func (f SaveSettingsArgumentFactory) New(accountID int, userID int, settingsInput []*model.AccountSettingInput) resolvers.Arguments[SaveSettingsArguments] {
	return SaveSettingsArguments{
		accountID: accountID,
		userID:    userID,
		settings:  settingsInput,
	}
}

func NewSaveSettingsValidator(userModel user.UserModel) resolvers.Validator[SaveSettingsArguments, ValidSaveSettingsArguments] {
	return validatorImpl{userModel: userModel}
}

func NewSaveSettingsResolver() resolvers.Resolver[model.Status, ValidSaveSettingsArguments] {
	return resolverImpl{}
}
