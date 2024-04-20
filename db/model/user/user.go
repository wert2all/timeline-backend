package user

import (
	"timeline/backend/db/repository/user"
	"timeline/backend/ent"
)

type SomeUser struct {
	UUID, Name, Email, Avatar string
}

type Authorize interface {
	CheckOrCreate(someUser SomeUser) (*ent.User, error)
}

type Get interface {
	GetByID(int) (ent.User, error)
}

type UserModel struct{ repository user.UserRepository }

func (u UserModel) CheckOrCreate(googleUser SomeUser) (*ent.User, error) {
	user, error := u.repository.FindByUUID(googleUser.UUID)
	error, notFound := error.(*ent.NotFoundError)

	if notFound {
		return u.repository.Create(googleUser.UUID, googleUser.Name, googleUser.Email, googleUser.Avatar)
	}

	if user.Active {
		return user, nil
	} else {
		return nil, error
	}
}

func (u UserModel) GetByID(userID int) (*ent.User, error) { return u.repository.FindByID(userID) }

func NewSomeUser(uuid, name, email, avatar string) SomeUser {
	return SomeUser{UUID: uuid, Name: name, Email: email, Avatar: avatar}
}

func NewUserModel(repository user.UserRepository) UserModel {
	return UserModel{
		repository: repository,
	}
}
