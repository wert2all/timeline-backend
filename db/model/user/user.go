package user

import (
	"timeline/backend/db/repository/user"
	"timeline/backend/ent"
)

type SomeUser struct {
	UUID, Name, Email, Avatar string
}

type Authorize interface {
	CheckOrCreate(someUser SomeUser) *ent.User
}

type UserModel struct {
	repository user.UserRepository
}

func (u UserModel) CheckOrCreate(googleUser SomeUser) *ent.User {
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

func NewSomeUser(uuid, name, email, avatar string) SomeUser {
	return SomeUser{UUID: uuid, Name: name, Email: email, Avatar: avatar}
}

func NewUserModel(repository user.UserRepository) UserModel {
	return UserModel{
		repository: repository,
	}
}
