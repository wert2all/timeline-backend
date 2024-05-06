package resolvers

import (
	"context"
	"errors"
	appContext "timeline/backend/app/context"
	"timeline/backend/db/model/timeline"
	"timeline/backend/db/model/user"
	"timeline/backend/graph/model"
)

type AddTimelineArguments struct {
	timeline *model.AddTimeline
}

func (a AddTimelineArguments) GetArguments() AddTimelineArguments { return a }

type addTimelimeMutation struct {
	Users    user.UserModel
	Timeline timeline.UserTimeline
}

func (a addTimelimeMutation) Resolve(ctx context.Context, arguments Arguments[AddTimelineArguments]) (*model.ShortUserTimeline, error) {
	userEntity, error := a.Users.GetUser(appContext.GetUserID(ctx))
	if error != nil {
		return nil, error
	}
	var name *string
	if arguments.GetArguments().timeline != nil {
		name = arguments.GetArguments().timeline.Name
	} else {
		return nil, errors.New(`missing required timeline`)
	}

	created, error := a.Timeline.CreateTimeline(name, userEntity)
	if error != nil {
		return nil, error
	}
	return &model.ShortUserTimeline{ID: created.ID, Name: &created.Name}, nil
}

func NewAddTimelineResolver(users user.UserModel, timeline timeline.UserTimeline) Resolver[*model.ShortUserTimeline, AddTimelineArguments] {
	return addTimelimeMutation{Users: users, Timeline: timeline}
}

func NewAddTimelineArguments(timeline *model.AddTimeline) Arguments[AddTimelineArguments] {
	return AddTimelineArguments{timeline: timeline}
}
