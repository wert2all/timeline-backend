package resolvers

import (
	"context"
	"errors"

	appContext "timeline/backend/app/context"
	"timeline/backend/db/model/timeline"
	"timeline/backend/db/model/user"
	"timeline/backend/ent"
	"timeline/backend/graph/model"
	"timeline/backend/lib/utils"

	"github.com/microcosm-cc/bluemonday"
)

type (
	AddTimelineArgumentFactory struct{}
	AddTimelineArguments       struct {
		timeline *model.AddTimeline
	}
	ValidAddTimelineArguments struct {
		name    string
		account *ent.Account
	}
	addTimelineValidator struct {
		UsersModel user.UserModel
	}
	addTimelimeMutation struct {
		Users    user.UserModel
		Timeline timeline.Timeline
	}
)

func (f AddTimelineArgumentFactory) New(timeline *model.AddTimeline) Arguments[AddTimelineArguments] {
	return AddTimelineArguments{timeline: timeline}
}

func (a addTimelineValidator) Validate(ctx context.Context, input Arguments[AddTimelineArguments]) (ValidArguments[ValidAddTimelineArguments], error) {
	p := bluemonday.StrictPolicy()
	account, error := a.UsersModel.GetUserAccount(input.GetArguments().timeline.AccountID, appContext.GetUserID(ctx))
	if error != nil {
		return nil, error
	}
	var name string
	if input.GetArguments().timeline != nil {
		name = utils.DerefString(input.GetArguments().timeline.Name)
	} else {
		return nil, errors.New(`missing required timeline`)
	}
	return ValidAddTimelineArguments{name: p.Sanitize(name), account: account}, nil
}
func (a AddTimelineArguments) GetArguments() AddTimelineArguments           { return a }
func (v ValidAddTimelineArguments) GetArguments() ValidAddTimelineArguments { return v }

func (a addTimelimeMutation) Resolve(_ context.Context, arguments ValidArguments[ValidAddTimelineArguments]) (*model.ShortUserTimeline, error) {
	created, error := a.Timeline.CreateTimeline(arguments.GetArguments().name, arguments.GetArguments().account)
	if error != nil {
		return nil, error
	}
	return &model.ShortUserTimeline{ID: created.ID, Name: &created.Name}, nil
}

func NewAddTimelineResolver(users user.UserModel, timeline timeline.Timeline) Resolver[*model.ShortUserTimeline, ValidAddTimelineArguments] {
	return addTimelimeMutation{Users: users, Timeline: timeline}
}

func NewAddtimelineValidator(userModel user.UserModel) Validator[AddTimelineArguments, ValidAddTimelineArguments] {
	return addTimelineValidator{UsersModel: userModel}
}
