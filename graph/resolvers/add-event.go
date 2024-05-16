package resolvers

import (
	"context"
	"net/url"
	"time"
	appContext "timeline/backend/app/context"
	"timeline/backend/db/model/event"
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
			SetShowTime(arguments.GetArguments().showTime).
			SetURL(arguments.GetArguments().url.String()))

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

	link, err := url.ParseRequestURI(utils.DerefString(arguments.GetArguments().eventInput.URL))
	if err != nil {
		panic(err)
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
		url:         link,
		tags:        maps.Values(groupedTags),
	}, err
}

func NewAddEventResolver(event event.Model, timeline timeline.UserTimeline) Resolver[*model.TimelineEvent, ValidAddEventArguments] {
	return addEventResolverImpl{event, timeline}
}

func NewAddEventValidator(timeline timeline.UserTimeline) Validator[AddEventArguments, ValidAddEventArguments] {
	return addEventvalidatorImpl{timeline}
}
