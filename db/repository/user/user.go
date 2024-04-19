package user

import (
	"fmt"
	"timeline/backend/ent"
	"timeline/backend/ent/user"

	"golang.org/x/net/context"
)

type UserRepository interface {
	FindByUUID(uuid string) *ent.User
	Create(uuid, name, email, avatar string) *ent.User
}

type UserRepositoryImpl struct {
	client  *ent.Client
	context context.Context
}

// Create implements UserRepository.
func (u UserRepositoryImpl) Create(uuid string, name string, email string, avatar string) *ent.User {
	user, err := u.client.User.Create().SetUUID(uuid).SetName(name).SetEmail(email).SetAvatar(avatar).Save(u.context)
	if err != nil {
		panic(fmt.Errorf("failed creating user: %w", err))
	}
	return user
}

// Create implements UserRepository.
func (u UserRepositoryImpl) FindByUUID(uuid string) *ent.User {
	user, err := u.client.User.Query().Where(user.UUID(uuid)).Only(u.context)
	if err != nil {
		return nil
	}
	return user
}

func NewUserRepository(ctx context.Context, client *ent.Client) UserRepository {
	return UserRepositoryImpl{
		client:  client,
		context: ctx,
	}
}
