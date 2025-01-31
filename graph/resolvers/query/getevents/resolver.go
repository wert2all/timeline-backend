package getevents

import (
	"context"

	"timeline/backend/db/model/event"
	"timeline/backend/db/model/tag"
	"timeline/backend/domain/db/cursor"
	"timeline/backend/ent"
	"timeline/backend/graph/convert"
	"timeline/backend/graph/model"
	"timeline/backend/graph/resolvers"

	"github.com/xorcare/pointer"
)

type (
	resolverImpl struct {
		eventModel event.Model
		tagModel   tag.Model
	}
)

func NewResolver(eventModel event.Model, tagModel tag.Model) resolvers.Resolver[*model.TimelineEvents, ValidGetCursorEventsArguments] {
	return resolverImpl{eventModel: eventModel, tagModel: tagModel}
}

func (r resolverImpl) Resolve(ctx context.Context, arguments resolvers.ValidArguments[ValidGetCursorEventsArguments]) (*model.TimelineEvents, error) {
	args := arguments.GetArguments()

	events, err := r.eventModel.GetTimelineEvents(args.timeline, args.accountID != nil, args.cursor, args.limit+1)
	if err != nil {
		return nil, err
	}

	events, lastEvent := r.modifyLastEvent(events, args.limit)

	return &model.TimelineEvents{
		Events: convert.ToEvents(events, r.createTags(events)),
		Page:   r.createPageInfo(args.cursor, lastEvent),
	}, nil
}

func (r resolverImpl) createTags(events []*ent.Event) map[int][]string {
	tags := make(map[int][]string)
	for _, event := range events {
		tagsEntities := r.tagModel.GetEventTags(event)
		entityTags := make([]string, 0)
		for _, tagEntity := range tagsEntities {
			entityTags = append(entityTags, tagEntity.Tag)
		}
		tags[event.ID] = entityTags
	}
	return tags
}

func (r resolverImpl) modifyLastEvent(events []*ent.Event, limit int) ([]*ent.Event, *ent.Event) {
	if len(events) == limit+1 {
		return events[:len(events)-1], events[limit]
	}
	return events, nil
}

func (r resolverImpl) createPageInfo(startCursor *cursor.Cursor, lastEvent *ent.Event) *model.PageInfo {
	hasNextPage := false
	var encodedStartCursor *string = nil
	var encodedEndCursor *string = nil
	var endCursor *cursor.Cursor = nil

	if lastEvent != nil {
		hasNextPage = true
		endCursor = cursor.NewCursor(lastEvent.ID, lastEvent.Date)
	}

	if startCursor != nil {
		encodedStartCursor = pointer.String(cursor.Encode(*startCursor))
	}

	if endCursor != nil {
		encodedEndCursor = pointer.String(cursor.Encode(*endCursor))
	}
	return &model.PageInfo{
		StartCursor: encodedStartCursor,
		EndCursor:   encodedEndCursor,
		HasNextPage: hasNextPage,
	}
}
