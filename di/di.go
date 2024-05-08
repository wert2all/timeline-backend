package di

import (
	"timeline/backend/db/model/event"
	"timeline/backend/db/model/timeline"
	"timeline/backend/db/model/user"
	"timeline/backend/graph/model"
	"timeline/backend/graph/resolvers"
)

type CORS struct {
	AllowedOrigin string
	Debug         bool
}
type Postgres struct {
	Host, User, Password, Database string
	Port                           int
}
type Config struct {
	Port          string
	CORS          CORS
	Postgres      Postgres
	GoogleClintID string
	SentryDsn     string
}

// Models todo refactor
type Models struct {
	Users    user.UserModel
	Timeline timeline.UserTimeline
	Event    event.Model
}
type Resolvers struct {
	MutationResolvers MutationResolvers
	QueryResolvers    QueryResolvers
}
type MutationResolvers struct {
	AddTimeline resolvers.Resolver[*model.ShortUserTimeline, resolvers.ValidAddTimelineArguments]
	Authorize   resolvers.Resolver[*model.User, resolvers.ValidAuthorizeArguments]
	AddEvent    resolvers.Resolver[*model.TimelineEvent, resolvers.ValidAddEventArguments]
}
type QueryResolvers struct{}

func NewAllModels(user user.UserModel, timeline timeline.UserTimeline, event event.Model) Models {
	return Models{Users: user, Timeline: timeline, Event: event}
}

func NewResolvers(event event.Model, users user.UserModel, timeline timeline.UserTimeline) Resolvers {
	return Resolvers{
		MutationResolvers: MutationResolvers{
			resolvers.NewAddTimelineResolver(users, timeline),
			resolvers.NewAutorizeResolver(timeline),
			resolvers.NewAddEventResolver(event, timeline),
		},
		QueryResolvers: QueryResolvers{},
	}
}
