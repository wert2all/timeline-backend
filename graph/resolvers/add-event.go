package resolvers

import (
	"context"
	"net/url"
	"time"

	"timeline/backend/db/model/event"
	"timeline/backend/db/model/tag"
	"timeline/backend/db/model/timeline"
	"timeline/backend/ent"
	entEvent "timeline/backend/ent/event"
	"timeline/backend/graph/model"
	eventValidator "timeline/backend/graph/resolvers/mutation/event"
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
	url         *url.URL
	tags        []string
}

type AddEventArguments struct {
	eventInput model.TimelineEventInput
}
type addEventResolverImpl struct {
	event    event.Model
	timeline timeline.UserTimeline
	tag      tag.Model
}

func (a AddEventArguments) GetArguments() AddEventArguments           { return a }
func (v ValidAddEventArguments) GetArguments() ValidAddEventArguments { return v }

func (a addEventResolverImpl) Resolve(ctx context.Context, arguments ValidArguments[ValidAddEventArguments]) (*model.TimelineEvent, error) {
	eventEntity, eventErr := a.event.Create(arguments.GetArguments().date, arguments.GetArguments().eventType)
	if eventErr != nil {
		return nil, eventErr
	}

	tags := make([]*ent.Tag, 0)
	for _, tagArgument := range arguments.GetArguments().tags {
		tagEntity, err := a.tag.UpsertTag(tagArgument)
		if err == nil {
			tags = append(tags, tagEntity)
		}
	}

	var shouldUpdateEntity *ent.EventUpdateOne
	shouldUpdateEntity = eventEntity.Update().
		SetTitle(arguments.GetArguments().title).
		SetDescription(arguments.GetArguments().description).
		SetShowTime(arguments.GetArguments().showTime).
		AddTags(tags...)

	if arguments.GetArguments().url != nil {
		shouldUpdateEntity = shouldUpdateEntity.SetURL(arguments.GetArguments().url.String())
	}

	updatedEntity, updateErr := a.event.Update(shouldUpdateEntity)

	if updateErr != nil {
		return nil, updateErr
	}

	_, err := a.timeline.AttachEvent(arguments.GetArguments().timeline, updatedEntity)
	if err != nil {
		return nil, err
	}

	tagEntities := a.tag.GetEventTags(updatedEntity)
	entityTags := make([]string, 0)

	for _, tagEntity := range tagEntities {
		entityTags = append(entityTags, tagEntity.Tag)
	}

	return &model.TimelineEvent{
		ID:          updatedEntity.ID,
		Date:        updatedEntity.Date,
		Type:        model.TimelineType(updatedEntity.Type.String()),
		Title:       &updatedEntity.Title,
		Description: &updatedEntity.Description,
		Tags:        entityTags,
	}, nil
}

type addEventvalidatorImpl struct {
	baseValidator eventValidator.BaseValidator
}

func (a addEventvalidatorImpl) Validate(ctx context.Context, arguments Arguments[AddEventArguments]) (ValidArguments[ValidAddEventArguments], error) {
	baseEvent, err := a.baseValidator.GetBaseValidEventInput(eventValidator.GQLInput(arguments.GetArguments().eventInput), ctx)
	if err != nil {
		return nil, err
	}
	return ValidAddEventArguments{
		timeline:    baseEvent.Timeline,
		eventType:   baseEvent.EventType,
		date:        baseEvent.Date,
		title:       baseEvent.Title,
		description: baseEvent.Description,
		showTime:    baseEvent.ShowTime,
		url:         baseEvent.Url,
		tags:        baseEvent.Tags,
	}, err
}

func NewAddEventResolver(event event.Model, timeline timeline.UserTimeline, tag tag.Model) Resolver[*model.TimelineEvent, ValidAddEventArguments] {
	return addEventResolverImpl{event, timeline, tag}
}

func NewAddEventValidator(timeline timeline.UserTimeline) Validator[AddEventArguments, ValidAddEventArguments] {
	return addEventvalidatorImpl{baseValidator: eventValidator.NewBaseValidator(timeline)}
}
