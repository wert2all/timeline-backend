package resolvers

import "context"
import "timeline/backend/graph/model"
import "timeline/backend/db/model/timeline"
import "timeline/backend/db/model/user"

type Arguments[T any] interface {
	GetArguments() T
}
type Resolver[T any, K any] interface {
	Resolve(context.Context, Arguments[K]) (T, error)
}
type MutationResolvers struct {
	AddTimeline Resolver[*model.ShortUserTimeline, AddTimelineArguments]
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
		},
		QueryResolvers: QueryResolvers{},
	}
}
