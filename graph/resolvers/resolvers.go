package resolvers

import "context"
import "timeline/backend/graph/model"
import "timeline/backend/db/model/timeline"
import "timeline/backend/db/model/user"

type Arguments[T any] interface {
	GetArguments() T
}

type ValidArguments[T any] interface {
	GetArguments() T
}

type Validator[K any, N any] interface {
	Validate(context.Context, Arguments[K]) (ValidArguments[N], error)
}
type Resolver[T any, K any] interface {
	Resolve(context.Context, ValidArguments[K]) (*T, error)
}

type MutationResolvers struct {
	AddTimeline Resolver[model.ShortUserTimeline, ValidAddTimelineArguments]
	Authorize   Resolver[model.User, ValidAuthorizeArguments]
}
type QueryResolvers struct{}
type Resolvers struct {
	MutationResolvers MutationResolvers
	QueryResolvers    QueryResolvers
}

func New(users user.UserModel, timeline timeline.UserTimeline) Resolvers {
	return Resolvers{
		MutationResolvers: MutationResolvers{
			NewAddTimelineResolver(users, timeline),
			NewAutorizeResolver(timeline),
		},
		QueryResolvers: QueryResolvers{},
	}
}

func Resolve[T any, K any, N any](context context.Context, inputArguments Arguments[K], validator Validator[K, N], resolver Resolver[T, N]) (*T, error) {
	valid, err := validator.Validate(context, inputArguments)
	if err != nil {
		return nil, err
	}
	return resolver.Resolve(context, valid)
}
