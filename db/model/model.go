package model

import (
	"timeline/backend/db/model/event"
	"timeline/backend/db/model/user"
)
import "timeline/backend/db/model/timeline"

// todo refactor
type AppModels struct {
	Users    user.UserModel
	Timeline timeline.UserTimeline
	Event    event.Model
}

func NewAllModels(user user.UserModel, timeline timeline.UserTimeline, event event.Model) AppModels {
	return AppModels{Users: user, Timeline: timeline, Event: event}
}
