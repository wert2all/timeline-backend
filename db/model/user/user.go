package user

import (
	"timeline/backend/db/repository/user"
	"timeline/backend/ent"
)

type GoogleUser struct {
	UUID, Name, Email, Avatar string
}

type Authorize interface {
	CheckOrCreate(googleUser GoogleUser) *ent.User
}

type UserModel struct {
	repository user.UserRepository
}

func (u UserModel) CheckOrCreate(googleUser GoogleUser) *ent.User {
	user := u.repository.FindByUUID(googleUser.UUID)

	if user == nil {
		return u.repository.Create(googleUser.UUID, googleUser.Name, googleUser.Email, googleUser.Avatar)
	}

	if user.Active {
		return user
	} else {
		return nil
	}
}

func NewUserModel(repository user.UserRepository) UserModel {
	return UserModel{
		repository: repository,
	}
}
