package user

import (
	"time"
	"timeline/backend/db/repository/user"
	"timeline/backend/ent"
)

type SomeUser struct {
	UUID, Name, Email, Avatar string
}

type CheckOrCreate struct {
	ID    int
	IsNew bool
}

type Authorize interface {
	CheckOrCreate(someUser SomeUser) (*CheckOrCreate, error)
	Authorize(int) (*ent.User, error)
}

type UserModel interface {
	Authorize
	GetUser(int) (*ent.User, error)
}

type userModelImp struct {
	repository user.UserRepository
}

func (u userModelImp) GetUser(userID int) (*ent.User, error) {
	return u.repository.FindByID(userID)
}

func (u userModelImp) CheckOrCreate(googleUser SomeUser) (*CheckOrCreate, error) {
	user, error := u.repository.FindByUUID(googleUser.UUID)
	error, notFound := error.(*ent.NotFoundError)

	if notFound {
		createdUser, error := u.repository.Create(googleUser.UUID, googleUser.Name, googleUser.Email, googleUser.Avatar)
		if error != nil {
			return nil, error
		}
		return &CheckOrCreate{ID: createdUser.ID, IsNew: true}, nil
	}

	if user.Active {
		return &CheckOrCreate{ID: user.ID, IsNew: false}, nil
	} else {
		return nil, error
	}
}

func (u userModelImp) Authorize(userID int) (*ent.User, error) {
	fetchedUser, error := u.repository.FindByID(userID)
	if error != nil {
		return nil, error
	}
	return u.repository.Save(fetchedUser.Update().SetUpdatedAt(time.Now()))
}

func NewSomeUser(uuid, name, email, avatar string) SomeUser {
	return SomeUser{UUID: uuid, Name: name, Email: email, Avatar: avatar}
}

func NewUserModel(repository user.UserRepository) UserModel {
	return userModelImp{
		repository: repository,
	}
}
