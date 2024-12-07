package resolvers

import (
	"context"

	appContext "timeline/backend/app/context"
	"timeline/backend/db/model/settings"
	"timeline/backend/db/model/timeline"
	"timeline/backend/db/model/user"
	"timeline/backend/ent"
	"timeline/backend/graph/model"
	enumvalues "timeline/backend/lib/enum-values"

	domainUser "timeline/backend/domain/user"
)

type ValidAuthorizeArguments struct {
	User  *ent.User
	IsNew bool
}

type authorizeResolverImpl struct {
	timeline timeline.Timeline
	users    user.UserModel
	settings settings.Model
}

type authorizeValidator struct {
	UsersModel    user.UserModel
	userExtractor domainUser.UserExtractor
}

type AuthorizeArgumentFactory struct{}

type AuthorizeArguments struct{}

func (v ValidAuthorizeArguments) GetArguments() ValidAuthorizeArguments { return v }
func (a AuthorizeArguments) GetArguments() AuthorizeArguments           { return a }
func (a authorizeValidator) Validate(ctx context.Context, _ Arguments[AuthorizeArguments]) (ValidArguments[ValidAuthorizeArguments], error) {
	token := appContext.GetToken(ctx)
	extractedUser, err := a.userExtractor.ExtractUserFromToken(ctx, &token)
	if err != nil {
		return nil, err
	}
	userEntity, err := a.UsersModel.GetUser(extractedUser.ID)
	if err != nil {
		return nil, err
	}

	return ValidAuthorizeArguments{User: userEntity, IsNew: extractedUser.IsNew}, nil
}

func (a authorizeResolverImpl) Resolve(_ context.Context, arguments ValidArguments[ValidAuthorizeArguments]) (*model.User, error) {
	userEntity := arguments.GetArguments().User
	accounts, err := a.users.GetUserAccounts(userEntity)
	if err != nil {
		return nil, err
	}

	gqlAccounts := make([]*model.ShortAccount, len(accounts))
	for i, accountEntity := range accounts {
		settings := a.settings.GetSettings(enumvalues.Account, accountEntity.ID)
		gqlSettings := make([]*model.AccountSettings, len(settings))

		for i, setting := range settings {
			gqlSettings[i] = &model.AccountSettings{
				Key:   setting.Key,
				Value: setting.Value,
			}
		}
		gqlAccounts[i] = &model.ShortAccount{
			Name:     &accountEntity.Name,
			Avatar:   accountEntity.Avatar,
			ID:       accountEntity.ID,
			Settings: gqlSettings,
		}
	}

	return &model.User{
		ID:       userEntity.ID,
		Name:     userEntity.Name,
		Email:    userEntity.Email,
		Avatar:   userEntity.Avatar,
		IsNew:    arguments.GetArguments().IsNew,
		Accounts: gqlAccounts,
	}, nil
}

func (a AuthorizeArgumentFactory) New() Arguments[AuthorizeArguments] { return AuthorizeArguments{} }

func NewAutorizeResolver(timeline timeline.Timeline, userModel user.UserModel, settings settings.Model) Resolver[*model.User, ValidAuthorizeArguments] {
	return authorizeResolverImpl{timeline: timeline, users: userModel, settings: settings}
}

func NewAuthorizeValidator(userModel user.UserModel, userExtractor domainUser.UserExtractor) Validator[AuthorizeArguments, ValidAuthorizeArguments] {
	return authorizeValidator{
		UsersModel:    userModel,
		userExtractor: userExtractor,
	}
}
