package di

import (
	"timeline/backend/db/model/event"
	"timeline/backend/db/model/tag"
	"timeline/backend/db/model/timeline"
	"timeline/backend/db/model/user"
)

type Postgres struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Config struct {
	App struct {
		Port string `yaml:"port"`
		Cors struct {
			Debug         bool   `yaml:"debug"`
			AllowedOrigin string `yaml:"allowedOrigin"`
		} `yaml:"cors"`
	} `yaml:"app"`

	Google struct {
		ClientId string `yaml:"clientId"`
	} `yaml:"google"`

	Postgres Postgres `yaml:"postgres"`

	Sentry struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"sentry"`
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
