package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	appContext "timeline/backend/app/context"
	"timeline/backend/graph/convert"
	"timeline/backend/graph/model"
)

// Authorize is the resolver for the authorize field.
func (r *mutationResolver) Authorize(ctx context.Context) (*model.User, error) {
	userEntity, error := r.Models.Users.Authorize(appContext.GetUserID(ctx))
	if error != nil {
		return nil, error
	}
	timelines, error := r.Models.Timeline.GetUserTimelines(userEntity)
	if error != nil {
		return nil, error
	}
	return convert.ToUser(userEntity, timelines, appContext.GetIsNew(ctx)), nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
