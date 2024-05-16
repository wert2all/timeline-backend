package di

import (
	"timeline/backend/db/model/event"
	"timeline/backend/db/model/tag"
	"timeline/backend/db/model/timeline"
	"timeline/backend/db/model/user"
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
	Tag      tag.Model
}

func NewAllModels(user user.UserModel, timeline timeline.UserTimeline, event event.Model, tag tag.Model) Models {
	return Models{Users: user, Timeline: timeline, Event: event, Tag: tag}
}
