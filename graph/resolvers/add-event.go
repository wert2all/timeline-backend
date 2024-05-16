package resolvers

import (
	"context"
	"net/url"
	"time"
	appContext "timeline/backend/app/context"
	"timeline/backend/db/model/event"
	"timeline/backend/db/model/tag"
	"timeline/backend/db/model/timeline"
	"timeline/backend/ent"
	entEvent "timeline/backend/ent/event"
	"timeline/backend/graph/model"
	"timeline/backend/lib/utils"

	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/exp/maps"
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

	tags := make([]*ent.Tag, len(arguments.GetArguments().tags))
	for i, tag := range arguments.GetArguments().tags {
		tagEntity, err := a.tag.UpsertTag(tag)
		if err == nil {
			tags[i] = tagEntity
		}
	}

	updatedEntity, updateErr := a.event.
		Update(eventEntity.Update().
			SetTitle(arguments.GetArguments().title).
			SetDescription(arguments.GetArguments().description).
			SetShowTime(arguments.GetArguments().showTime).
			SetURL(arguments.GetArguments().url.String()).
			AddTags(tags...))

	if updateErr != nil {
		return nil, updateErr
	}

	_, err := a.timeline.AttachEvent(arguments.GetArguments().timeline, updatedEntity)
	if err != nil {
		return nil, err
	}

	tagEntities := a.tag.GetEventTags(updatedEntity)
	entityTags := make([]string, len(tagEntities))

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
	Timeline timeline.UserTimeline
}

func (a addEventvalidatorImpl) Validate(ctx context.Context, arguments Arguments[AddEventArguments]) (ValidArguments[ValidAddEventArguments], error) {
	p := bluemonday.StrictPolicy()
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

	groupedTags := make(map[string]string)
	for _, tag := range input.Tags {
		groupedTags[p.Sanitize(tag)] = p.Sanitize(tag)
	}

	return ValidAddEventArguments{
		timeline:    timelineEntity,
		eventType:   eventType,
		date:        arguments.GetArguments().eventInput.Date,
		title:       p.Sanitize(utils.DerefString(arguments.GetArguments().eventInput.Title)),
		description: p.Sanitize(utils.DerefString(arguments.GetArguments().eventInput.Description)),
		showTime:    *input.ShowTime,
		url:         a.getLink(arguments.GetArguments().eventInput.URL),
		tags:        maps.Values(groupedTags),
	}, err
}

func (a addEventvalidatorImpl) getLink(link *string) *url.URL {
	linkString := utils.DerefString(link)
	if linkString != "" {
		link, err := url.ParseRequestURI(linkString)
		if err != nil {
			panic(err)
		}
		return link
	} else {
		return nil
	}
}

func NewAddEventResolver(event event.Model, timeline timeline.UserTimeline, tag tag.Model) Resolver[*model.TimelineEvent, ValidAddEventArguments] {
	return addEventResolverImpl{event, timeline, tag}
}

func NewAddEventValidator(timeline timeline.UserTimeline) Validator[AddEventArguments, ValidAddEventArguments] {
	return addEventvalidatorImpl{timeline}
}
