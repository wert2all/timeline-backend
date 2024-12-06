package settings

import (
	"context"

	"timeline/backend/graph/model"
	"timeline/backend/graph/resolvers"
)

type (
	SaveSettingsArguments struct {
		accountID int
		userID    int
		settings  []*model.AccountSettingInput
	}
	ValidSaveSettingsArguments  struct{}
	SaveSettingsArgumentFactory struct{}

	validatorImpl struct{}
	resolverImpl  struct{}
)

// GetArguments implements resolvers.Arguments.
func (s SaveSettingsArguments) GetArguments() SaveSettingsArguments { return s }

// Resolve implements resolvers.Resolver.
func (r resolverImpl) Resolve(ctx context.Context, arguments resolvers.ValidArguments[ValidSaveSettingsArguments]) (model.Status, error) {
	panic("unimplemented")
}

// Validate implements resolvers.Validator.
func (v validatorImpl) Validate(ctx context.Context, arguments resolvers.Arguments[SaveSettingsArguments]) (resolvers.ValidArguments[ValidSaveSettingsArguments], error) {
	panic("unimplemented")
}

func (f SaveSettingsArgumentFactory) New(accountID int, userID int, settingsInput []*model.AccountSettingInput) resolvers.Arguments[SaveSettingsArguments] {
	return SaveSettingsArguments{
		accountID: accountID,
		userID:    userID,
		settings:  settingsInput,
	}
}

func NewSaveSettingsValidator() resolvers.Validator[SaveSettingsArguments, ValidSaveSettingsArguments] {
	return validatorImpl{}
}

func NewSaveSettingsResolver() resolvers.Resolver[model.Status, ValidSaveSettingsArguments] {
	return resolverImpl{}
}
