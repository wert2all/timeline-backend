package user

import (
	"fmt"

	"timeline/backend/ent"
	"timeline/backend/ent/account"
	"timeline/backend/ent/user"

	"golang.org/x/net/context"
)

type Repository interface {
	FindByID(ID int) (*ent.User, error)
	FindByUUID(uuid string) (*ent.User, error)
	Create(uuid, name, email, avatar string) (*ent.User, error)
	Save(*ent.UserUpdateOne) (*ent.User, error)
	AddAccount(*ent.User, *ent.Account) (*ent.User, error)
	GetUserAccounts(*ent.User) ([]*ent.Account, error)
	GetUserAccount(int, int) (*ent.Account, error)
}

type userRepositoryImpl struct {
	client  *ent.Client
	context context.Context
}

// GetUserAccount implements Repository.
func (u userRepositoryImpl) GetUserAccount(accountID int, userID int) (*ent.Account, error) {
	return u.client.Account.
		Query().
		Where(
			account.ID(accountID),
			account.HasUserWith(user.ID(userID)),
		).
		Only(u.context)
}

// GetUserAccounts implements Repository.
func (u userRepositoryImpl) GetUserAccounts(user *ent.User) ([]*ent.Account, error) {
	return user.QueryAccount().All(u.context)
}

// AddAccount implements Repository.
func (u userRepositoryImpl) AddAccount(user *ent.User, account *ent.Account) (*ent.User, error) {
	return user.Update().AddAccount(account).Save(u.context)
}

// Save implements  Repository.
func (u userRepositoryImpl) Save(user *ent.UserUpdateOne) (*ent.User, error) {
	return user.Save(u.context)
}

// FindByID implements  Repository.
func (u userRepositoryImpl) FindByID(ID int) (*ent.User, error) {
	return u.client.User.Query().Where(user.ID(ID)).Only(u.context)
}

// Create implements  Repository.
func (u userRepositoryImpl) Create(uuid string, name string, email string, avatar string) (*ent.User, error) {
	userEntity, err := u.client.User.Create().SetUUID(uuid).SetName(name).SetEmail(email).SetAvatar(avatar).Save(u.context)
	if err != nil {
		panic(fmt.Errorf("failed creating userEntity: %w", err))
	}
	return userEntity, nil
}

// FindByUUID implements Repository.
func (u userRepositoryImpl) FindByUUID(uuid string) (*ent.User, error) {
	return u.client.User.Query().Where(user.UUID(uuid)).Only(u.context)
}

func NewUserRepository(ctx context.Context, client *ent.Client) Repository {
	return userRepositoryImpl{
		client:  client,
		context: ctx,
	}
}
