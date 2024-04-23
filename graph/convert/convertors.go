package convert

import (
	"timeline/backend/ent"
	"timeline/backend/graph/model"
)

func ToUser(user *ent.User, isNew bool) *model.User {
	return &model.User{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Avatar: user.Avatar,
		IsNew:  isNew,
	}
}
