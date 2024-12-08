package settings

import (
	"context"
	"errors"

	appContext "timeline/backend/app/context"
	"timeline/backend/db/model/settings"
	"timeline/backend/db/model/user"
	domainUser "timeline/backend/domain/user"
	"timeline/backend/graph/model"
	"timeline/backend/graph/resolvers"
	enumvalues "timeline/backend/lib/enum-values"
)

type (
	SaveSettingsArguments struct {
		accountID int
		settings  []*model.AccountSettingInput
	}
	ValidSaveSettingsArguments struct {
		accountID int
		settings  map[string]string
	}
	SaveSettingsArgumentFactory struct{}

	validatorImpl struct {
		userModel     user.UserModel
		userExtractor domainUser.UserExtractor
	}
	resolverImpl struct {
		settingsModel settings.Model
	}
)

// GetArguments implements resolvers.ValidArguments.
func (v ValidSaveSettingsArguments) GetArguments() ValidSaveSettingsArguments { return v }

// GetArguments implements resolvers.Arguments.
func (s SaveSettingsArguments) GetArguments() SaveSettingsArguments { return s }

// Resolve implements resolvers.Resolver.
func (r resolverImpl) Resolve(ctx context.Context, arguments resolvers.ValidArguments[ValidSaveSettingsArguments]) (model.Status, error) {
	_, err := r.settingsModel.SaveSettings(enumvalues.SettingsTypeAccount, arguments.GetArguments().accountID, arguments.GetArguments().settings)
	if err != nil {
		return model.StatusError, errors.New("could not save settings")
	} else {
		return model.StatusSuccess, nil
	}
}

// Validate implements resolvers.Validator.
func (v validatorImpl) Validate(ctx context.Context, arguments resolvers.Arguments[SaveSettingsArguments]) (resolvers.ValidArguments[ValidSaveSettingsArguments], error) {
	token := appContext.GetToken(ctx)
	user, errExtraction := v.userExtractor.ExtractUserFromToken(ctx, &token)
	if errExtraction != nil {
		return nil, errors.New("could not save settings")
	}
	args := arguments.GetArguments()
	_, err := v.userModel.GetUserAccount(args.accountID, user.ID)
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

func (f SaveSettingsArgumentFactory) New(accountID int, settingsInput []*model.AccountSettingInput) resolvers.Arguments[SaveSettingsArguments] {
	return SaveSettingsArguments{
		accountID: accountID,
		settings:  settingsInput,
	}
}

func NewSaveSettingsValidator(userModel user.UserModel, userExtractor domainUser.UserExtractor) resolvers.Validator[SaveSettingsArguments, ValidSaveSettingsArguments] {
	return validatorImpl{userModel: userModel, userExtractor: userExtractor}
}

func NewSaveSettingsResolver(settingsModel settings.Model) resolvers.Resolver[model.Status, ValidSaveSettingsArguments] {
	return resolverImpl{settingsModel: settingsModel}
}
