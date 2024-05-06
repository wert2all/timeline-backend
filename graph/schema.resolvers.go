package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

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
		resolvers.NewAuthorizeArguments(),
		resolvers.NewAuthorizeValidator(r.Models.Users),
		r.Resolvers.MutationResolvers.Authorize,
	)
}

// AddTimeline is the resolver for the addTimeline field.
func (r *mutationResolver) AddTimeline(ctx context.Context, timeline *model.AddTimeline) (*model.ShortUserTimeline, error) {
	return resolvers.Resolve(
		ctx,
		resolvers.NewAddTimelineArguments(timeline),
		resolvers.NewAddtimelineValidator(),
		r.Resolvers.MutationResolvers.AddTimeline,
	)
}

// AddEvent is the resolver for the addEvent field.
func (r *mutationResolver) AddEvent(ctx context.Context, event model.TimelineEventInput) (*model.TimelineEvent, error) {
	return resolvers.Resolve(
		ctx,
		resolvers.NewAddEventArguments(event),
		resolvers.NewAddEventValidator(r.Models.Timeline),
		r.Resolvers.MutationResolvers.AddEvent,
	)
}

// TimelineEvents is the resolver for the timelineEvents field.
func (r *queryResolver) TimelineEvents(ctx context.Context, timelineID int, limit *model.Limit) ([]*model.TimelineEvent, error) {
	timeline, error := r.Models.Timeline.GetUserTimeline(appContext.GetUserID(ctx), timelineID)
	if error != nil {
		return nil, error
	}
	events, error := r.Models.Timeline.GetEvents(timeline, convert.ToLimit(limit))
	if error != nil {
		return nil, error
	}
	return convert.ToEvents(events), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
