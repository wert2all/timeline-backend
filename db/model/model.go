package model

import "timeline/backend/db/model/user"

// todo refactor
type AppModels struct {
	Users user.UserModel
}

func NewAllModels(user user.UserModel) AppModels { return AppModels{Users: user} }
