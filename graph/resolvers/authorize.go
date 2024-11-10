package resolvers

import (
	"context"

	appContext "timeline/backend/app/context"
	"timeline/backend/db/model/timeline"
	"timeline/backend/db/model/user"
	"timeline/backend/ent"
	"timeline/backend/graph/model"
)

type ValidAuthorizeArguments struct {
	User  *ent.User
	IsNew bool
}

type authorizeResolverImpl struct {
	timeline timeline.Timeline
	users    user.UserModel
}

type authorizeValidator struct {
	UsersModel user.UserModel
}

type AuthorizeArgumentFactory struct{}

type AuthorizeArguments struct{}

func (v ValidAuthorizeArguments) GetArguments() ValidAuthorizeArguments { return v }
func (a AuthorizeArguments) GetArguments() AuthorizeArguments           { return a }
func (a authorizeValidator) Validate(ctx context.Context, _ Arguments[AuthorizeArguments]) (ValidArguments[ValidAuthorizeArguments], error) {
	userEntity, err := a.UsersModel.GetUser(appContext.GetUserID(ctx))
	if err != nil {
		return nil, err
	}

	return ValidAuthorizeArguments{User: userEntity, IsNew: appContext.GetIsNew(ctx)}, nil
}

func (a authorizeResolverImpl) Resolve(_ context.Context, arguments ValidArguments[ValidAuthorizeArguments]) (*model.User, error) {
	userEntity := arguments.GetArguments().User
	accounts, err := a.users.GetUserAccounts(userEntity)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:       userEntity.ID,
		Name:     userEntity.Name,
		Email:    userEntity.Email,
		Avatar:   userEntity.Avatar,
		IsNew:    arguments.GetArguments().IsNew,
		Accounts: a.convertAccounts(accounts),
	}, nil
}

func (a AuthorizeArgumentFactory) New() Arguments[AuthorizeArguments] { return AuthorizeArguments{} }

func NewAutorizeResolver(timeline timeline.Timeline, userModel user.UserModel) Resolver[*model.User, ValidAuthorizeArguments] {
	return authorizeResolverImpl{timeline: timeline, users: userModel}
}

func NewAuthorizeValidator(userModel user.UserModel) Validator[AuthorizeArguments, ValidAuthorizeArguments] {
	return authorizeValidator{UsersModel: userModel}
}

// func (a authorizeResolverImpl) converTimelines(timelines []*ent.Timeline) []*model.ShortUserTimeline {
// 	gqlTimelines := make([]*model.ShortUserTimeline, len(timelines))
// 	for i, timelineEntity := range timelines {
// 		gqlTimelines[i] = &model.ShortUserTimeline{ID: timelineEntity.ID, Name: &timelineEntity.Name}
// 	}
// 	return gqlTimelines
// }

func (a authorizeResolverImpl) convertAccounts(accounts []*ent.Account) []*model.ShortAccount {
	gqlAccounts := make([]*model.ShortAccount, len(accounts))
	for i, accountEntity := range accounts {
		gqlAccounts[i] = &model.ShortAccount{Name: &accountEntity.Name, Avatar: accountEntity.Avatar, ID: accountEntity.ID}
	}
	return gqlAccounts
}
