package resolvers

import (
	"context"
	"time"
	appContext "timeline/backend/app/context"
	"timeline/backend/db/model/event"
	"timeline/backend/db/model/timeline"
	"timeline/backend/ent"
	entEvent "timeline/backend/ent/event"
	"timeline/backend/graph/model"
)

type AddEventArgumentFactory struct{}

func (f AddEventArgumentFactory) New(input model.TimelineEventInput) Arguments[AddEventArguments] {
	return AddEventArguments{eventInput: input}
}

type ValidAddEventArguments struct {
	timeline    *ent.Timeline
	eventType   entEvent.Type
	date        time.Time
	title       string
	description string
	showTime    bool
}

type AddEventArguments struct {
	eventInput model.TimelineEventInput
}
type addEventResolverImpl struct {
	event    event.Model
	timeline timeline.UserTimeline
}

func (a AddEventArguments) GetArguments() AddEventArguments           { return a }
func (v ValidAddEventArguments) GetArguments() ValidAddEventArguments { return v }

func (a addEventResolverImpl) Resolve(ctx context.Context, arguments ValidArguments[ValidAddEventArguments]) (*model.TimelineEvent, error) {
	eventEntity, eventErr := a.event.Create(arguments.GetArguments().date, arguments.GetArguments().eventType)
	if eventErr != nil {
		return nil, eventErr
	}

	updatedEntity, updateErr := a.event.
		Update(eventEntity.Update().
			SetTitle(arguments.GetArguments().title).
			SetDescription(arguments.GetArguments().description).
			SetShowTime(arguments.GetArguments().showTime))

	if updateErr != nil {
		return nil, updateErr
	}

	_, err := a.timeline.AttachEvent(arguments.GetArguments().timeline, updatedEntity)
	if err != nil {
		return nil, err
	}

	return &model.TimelineEvent{
		ID:          updatedEntity.ID,
		Date:        updatedEntity.Date,
		Type:        model.TimelineType(updatedEntity.Type.String()),
		Title:       &updatedEntity.Title,
		Description: &updatedEntity.Description,
	}, nil
}

type addEventvalidatorImpl struct {
	Timeline timeline.UserTimeline
}

func (a addEventvalidatorImpl) Validate(ctx context.Context, arguments Arguments[AddEventArguments]) (ValidArguments[ValidAddEventArguments], error) {
	input := arguments.GetArguments().eventInput
	timelineEntity, err := a.Timeline.GetUserTimeline(appContext.GetUserID(ctx), input.TimelineID)
	if err != nil {
		return nil, err
	}
	var eventType entEvent.Type
	if input.Type == nil {
		eventType = entEvent.Type(model.TimelineTypeDefault)
	} else {
		eventType = entEvent.Type(input.Type.String())
	}

	return ValidAddEventArguments{
		timeline:    timelineEntity,
		eventType:   eventType,
		date:        arguments.GetArguments().eventInput.Date,
		title:       derefString(arguments.GetArguments().eventInput.Title),
		description: derefString(arguments.GetArguments().eventInput.Description),
		showTime:    *input.ShowTime,
	}, err
}

func derefString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func NewAddEventResolver(event event.Model, timeline timeline.UserTimeline) Resolver[*model.TimelineEvent, ValidAddEventArguments] {
	return addEventResolverImpl{event, timeline}
}

func NewAddEventValidator(timeline timeline.UserTimeline) Validator[AddEventArguments, ValidAddEventArguments] {
	return addEventvalidatorImpl{timeline}
}
