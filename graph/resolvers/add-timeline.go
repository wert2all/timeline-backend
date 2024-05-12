package resolvers

import (
	"context"
	"errors"
	"github.com/microcosm-cc/bluemonday"
	appContext "timeline/backend/app/context"
	"timeline/backend/db/model/timeline"
	"timeline/backend/db/model/user"
	"timeline/backend/ent"
	"timeline/backend/graph/model"
	"timeline/backend/lib/utils"
)

type AddTimelineArgumentFactory struct{}

func (f AddTimelineArgumentFactory) New(timeline *model.AddTimeline) Arguments[AddTimelineArguments] {
	return AddTimelineArguments{timeline: timeline}
}

type AddTimelineArguments struct {
	timeline *model.AddTimeline
}

type ValidAddTimelineArguments struct {
	name string
	User *ent.User
}

type addTimelineValidator struct {
	UsersModel user.UserModel
}

type addTimelimeMutation struct {
	Users    user.UserModel
	Timeline timeline.UserTimeline
}

func (a addTimelineValidator) Validate(ctx context.Context, input Arguments[AddTimelineArguments]) (ValidArguments[ValidAddTimelineArguments], error) {
	p := bluemonday.StrictPolicy()
	userEntity, error := a.UsersModel.GetUser(appContext.GetUserID(ctx))
	if error != nil {
		return nil, error
	}
	var name string
	if input.GetArguments().timeline != nil {
		name = utils.DerefString(input.GetArguments().timeline.Name)
	} else {
		return nil, errors.New(`missing required timeline`)
	}
	return ValidAddTimelineArguments{name: p.Sanitize(name), User: userEntity}, nil
}
func (a AddTimelineArguments) GetArguments() AddTimelineArguments           { return a }
func (v ValidAddTimelineArguments) GetArguments() ValidAddTimelineArguments { return v }

func (a addTimelimeMutation) Resolve(_ context.Context, arguments ValidArguments[ValidAddTimelineArguments]) (*model.ShortUserTimeline, error) {
	created, error := a.Timeline.CreateTimeline(arguments.GetArguments().name, arguments.GetArguments().User)
	if error != nil {
		return nil, error
	}
	return &model.ShortUserTimeline{ID: created.ID, Name: &created.Name}, nil
}

func NewAddTimelineResolver(users user.UserModel, timeline timeline.UserTimeline) Resolver[*model.ShortUserTimeline, ValidAddTimelineArguments] {
	return addTimelimeMutation{Users: users, Timeline: timeline}
}

func NewAddtimelineValidator(userModel user.UserModel) Validator[AddTimelineArguments, ValidAddTimelineArguments] {
	return addTimelineValidator{UsersModel: userModel}
}
