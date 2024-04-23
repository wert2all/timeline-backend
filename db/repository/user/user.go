package user

import (
	"fmt"
	"timeline/backend/ent"
	"timeline/backend/ent/user"

	"golang.org/x/net/context"
)

type UserRepository interface {
	FindByID(ID int) (*ent.User, error)
	FindByUUID(uuid string) (*ent.User, error)
	Create(uuid, name, email, avatar string) (*ent.User, error)
	Save(user *ent.UserUpdateOne) (*ent.User, error)
}

type userRepositoryImpl struct {
	client  *ent.Client
	context context.Context
}

// Save implements UserRepository.
func (u userRepositoryImpl) Save(user *ent.UserUpdateOne) (*ent.User, error) {
	return user.Save(u.context)
}

// FindByID implements UserRepository.
func (u userRepositoryImpl) FindByID(ID int) (*ent.User, error) {
	return u.client.User.Query().Where(user.ID(ID)).Only(u.context)
}

// Create implements UserRepository.
func (u userRepositoryImpl) Create(uuid string, name string, email string, avatar string) (*ent.User, error) {
	user, err := u.client.User.Create().SetUUID(uuid).SetName(name).SetEmail(email).SetAvatar(avatar).Save(u.context)
	if err != nil {
		panic(fmt.Errorf("failed creating user: %w", err))
	}
	return user, nil
}

// Create implements UserRepository.
func (u userRepositoryImpl) FindByUUID(uuid string) (*ent.User, error) {
	return u.client.User.Query().Where(user.UUID(uuid)).Only(u.context)
}

func NewUserRepository(ctx context.Context, client *ent.Client) UserRepository {
	return userRepositoryImpl{
		client:  client,
		context: ctx,
	}
}
