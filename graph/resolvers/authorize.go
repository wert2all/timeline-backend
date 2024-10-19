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
	Timeline timeline.UserTimeline
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
	timelines, err := a.Timeline.GetUserTimelines(arguments.GetArguments().User)
	if err != nil {
		return nil, err
	}
	userEntity := arguments.GetArguments().User
	return &model.User{
		ID:        userEntity.ID,
		Name:      userEntity.Name,
		Email:     userEntity.Email,
		Avatar:    userEntity.Avatar,
		IsNew:     arguments.GetArguments().IsNew,
		Timelines: converTimelines(timelines),
	}, nil
}
func (a AuthorizeArgumentFactory) New() Arguments[AuthorizeArguments] { return AuthorizeArguments{} }

func NewAutorizeResolver(timeline timeline.UserTimeline) Resolver[*model.User, ValidAuthorizeArguments] {
	return authorizeResolverImpl{Timeline: timeline}
}

func NewAuthorizeValidator(userModel user.UserModel) Validator[AuthorizeArguments, ValidAuthorizeArguments] {
	return authorizeValidator{UsersModel: userModel}
}

func converTimelines(timelines []*ent.Timeline) []*model.ShortUserTimeline {
	gqlTimelines := make([]*model.ShortUserTimeline, len(timelines))
	for i, timelineEntity := range timelines {
		gqlTimelines[i] = &model.ShortUserTimeline{ID: timelineEntity.ID, Name: &timelineEntity.Name}
	}
	return gqlTimelines
}
