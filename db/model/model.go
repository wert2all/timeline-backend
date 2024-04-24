package model

import "timeline/backend/db/model/user"
import "timeline/backend/db/model/timeline"

// todo refactor
type AppModels struct {
	Users    user.UserModel
	Timeline timeline.UserTimeline
}

func NewAllModels(user user.UserModel, timeline timeline.UserTimeline) AppModels {
	return AppModels{Users: user, Timeline: timeline}
}
