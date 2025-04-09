package getevent

import (
	"context"
	"errors"

	"timeline/backend/db/model/event"
	"timeline/backend/db/model/tag"
	"timeline/backend/ent"
	"timeline/backend/graph/convert"
	"timeline/backend/graph/model"
	"timeline/backend/graph/resolvers"
)

type (
	resolverImpl struct {
		eventModel event.Model
		tagModel   tag.Model
	}
)

// Resolve implements resolvers.Resolver.
func (r resolverImpl) Resolve(ctx context.Context, arguments resolvers.ValidArguments[ValidGetEventArguments]) (*model.TimelineEvent, error) {
	args := arguments.GetArguments()
	event, err := r.eventModel.GetEventByID(args.eventID)
	if err != nil {
		return nil, err
	}

	timeline, err := event.QueryTimeline().Only(ctx)
	if err != nil {
		return nil, err
	}
	if err := r.canShow(ctx, event, arguments, timeline); err != nil {
		return nil, err
	}

	return convert.ToEvent(event, r.getTags(event), timeline.ID), nil
}

func (r resolverImpl) getTags(event *ent.Event) []string {
	entityTags := make([]string, 0)

	tagsEntities := r.tagModel.GetEventTags(event)
	for _, tagEntity := range tagsEntities {
		entityTags = append(entityTags, tagEntity.Tag)
	}
	return entityTags
}

func (r resolverImpl) canShow(ctx context.Context, event *ent.Event, arguments resolvers.ValidArguments[ValidGetEventArguments], timeline *ent.Timeline) error {
	args := arguments.GetArguments()
	if event.Private {
		if args.account != nil {
			accountID, err := timeline.QueryAccount().OnlyID(ctx)
			if err != nil {
				return err
			}

			if accountID != args.account.ID {
				return errors.New("not found")
			}

		} else {
			return errors.New("not found")
		}
	}
	return nil
}

func NewResolver(eventModel event.Model, tagModel tag.Model) resolvers.Resolver[*model.TimelineEvent, ValidGetEventArguments] {
	return resolverImpl{eventModel: eventModel, tagModel: tagModel}
}
