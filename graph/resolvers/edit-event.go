package resolvers

import (
	"context"

	"timeline/backend/db/model/event"
	"timeline/backend/db/model/timeline"
	"timeline/backend/graph/model"
	eventValidator "timeline/backend/graph/resolvers/mutation/event"
)

type (
	EditEventArgumentFactory struct{}

	EditEventArguments struct {
		eventInput model.ExistTimelineEventInput
	}

	ValidEditEventArguments struct {
		id        int
		baseInput *eventValidator.BaseValidEventInput
	}

	editEventvalidatorImpl struct {
		baseValidator eventValidator.BaseValidator
		eventsModel   event.Model
		timelineModel timeline.Timeline
	}

	editEventResolverImpl struct {
		eventsModel event.Model
	}
)

// Resolve implements Resolver.
func (e editEventResolverImpl) Resolve(ctx context.Context, arguments ValidArguments[ValidEditEventArguments]) (*model.TimelineEvent, error) {
	eventEntity, err := e.eventsModel.GetEventByID(arguments.GetArguments().id)
	if err != nil {
		return nil, err
	}
	updatedEntity, updateErr := e.eventsModel.UpdateEvent(eventEntity, arguments.GetArguments().baseInput)

	if updateErr != nil {
		return nil, updateErr
	}

	return &model.TimelineEvent{
		ID:          updatedEntity.ID,
		Date:        updatedEntity.Date,
		Type:        model.TimelineType(updatedEntity.Type.String()),
		Title:       &updatedEntity.Title,
		Description: &updatedEntity.Description,
		Tags:        arguments.GetArguments().baseInput.Tags,
	}, nil
}

// Validate implements Validator.
func (e editEventvalidatorImpl) Validate(ctx context.Context, arguments Arguments[EditEventArguments]) (ValidArguments[ValidEditEventArguments], error) {
	input := arguments.GetArguments().eventInput
	gqlInput := eventValidator.GQLInput{
		TimelineID:       input.TimelineID,
		Type:             input.Type,
		Date:             input.Date,
		Title:            input.Title,
		Description:      input.Description,
		ShowTime:         input.ShowTime,
		URL:              input.URL,
		Tags:             input.Tags,
		PreviewlyImageID: input.PreviewlyImageID,
	}
	baseEvent, err := e.baseValidator.GetBaseValidEventInput(gqlInput, ctx)
	if err != nil {
		return nil, err
	}

	return ValidEditEventArguments{
		id:        input.ID,
		baseInput: baseEvent,
	}, err
}

func (v ValidEditEventArguments) GetArguments() ValidEditEventArguments { return v }
func (e EditEventArgumentFactory) New(input model.ExistTimelineEventInput) Arguments[EditEventArguments] {
	return EditEventArguments{eventInput: input}
}

// GetArguments implements Arguments.
func (e EditEventArguments) GetArguments() EditEventArguments { return e }

func NewEditEventValidator(baseValidator eventValidator.BaseValidator, eventModel event.Model, timelineModel timeline.Timeline) Validator[EditEventArguments, ValidEditEventArguments] {
	return editEventvalidatorImpl{baseValidator: baseValidator, eventsModel: eventModel, timelineModel: timelineModel}
}

func NewEditEventResolver(eventModel event.Model) Resolver[*model.TimelineEvent, ValidEditEventArguments] {
	return editEventResolverImpl{eventsModel: eventModel}
}
