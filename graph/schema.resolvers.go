package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	appContext "timeline/backend/app/context"
	"timeline/backend/graph/convert"
	"timeline/backend/graph/model"
	"timeline/backend/graph/resolvers"
)

// Authorize is the resolver for the authorize field.
func (r *mutationResolver) Authorize(ctx context.Context) (*model.User, error) {
	return resolvers.Resolve(
		ctx,
		r.ServiceLocator.Resolvers().Mutation().Authorize().ArgumentFactory().New(),
		r.ServiceLocator.Resolvers().Mutation().Authorize().Validator(),
		r.ServiceLocator.Resolvers().Mutation().Authorize().Resolver(),
	)
}

// AddTimeline is the resolver for the addTimeline field.
func (r *mutationResolver) AddTimeline(ctx context.Context, timeline *model.AddTimeline) (*model.ShortUserTimeline, error) {
	return resolvers.Resolve(
		ctx,
		r.ServiceLocator.Resolvers().Mutation().AddTimeline().ArgumentFactory().New(timeline),
		r.ServiceLocator.Resolvers().Mutation().AddTimeline().Validator(),
		r.ServiceLocator.Resolvers().Mutation().AddTimeline().Resolver(),
	)
}

// AddEvent is the resolver for the addEvent field.
func (r *mutationResolver) AddEvent(ctx context.Context, event model.TimelineEventInput) (*model.TimelineEvent, error) {
	return resolvers.Resolve(
		ctx,
		r.ServiceLocator.Resolvers().Mutation().AddEvent().ArgumentFactory().New(event),
		r.ServiceLocator.Resolvers().Mutation().AddEvent().Validator(),
		r.ServiceLocator.Resolvers().Mutation().AddEvent().Resolver(),
	)
}

// DeleteEvent is the resolver for the deleteEvent field.
func (r *mutationResolver) DeleteEvent(ctx context.Context, eventID int) (model.Status, error) {
	return resolvers.Resolve(
		ctx,
		r.ServiceLocator.Resolvers().Mutation().DeleteEvent().ArgumentFactory().New(eventID),
		r.ServiceLocator.Resolvers().Mutation().DeleteEvent().Validator(),
		r.ServiceLocator.Resolvers().Mutation().DeleteEvent().Resolver(),
	)
}

// TimelineEvents is the resolver for the timelineEvents field.
func (r *queryResolver) TimelineEvents(ctx context.Context, timelineID int, limit *model.Limit) ([]*model.TimelineEvent, error) {
	timeline, error := r.ServiceLocator.Models().Timeline().GetUserTimeline(appContext.GetUserID(ctx), timelineID)
	if error != nil {
		return nil, error
	}
	events, error := r.ServiceLocator.Models().Timeline().GetEvents(timeline, convert.ToLimit(limit))
	if error != nil {
		return nil, error
	}

	tags := make(map[int][]string)
	for _, event := range events {
		tagsEntities := r.ServiceLocator.Models().Tag().GetEventTags(event)
		entityTags := make([]string, 0)
		for _, tagEntity := range tagsEntities {
			entityTags = append(entityTags, tagEntity.Tag)
		}
		tags[event.ID] = entityTags
	}

	return convert.ToEvents(events, tags), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
