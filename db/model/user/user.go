package user

import (
	"context"
	"time"

	"timeline/backend/db/repository/account"
	"timeline/backend/db/repository/user"
	"timeline/backend/ent"
	"timeline/backend/graph/client/previewly"
)

type SomeUser struct {
	UUID, Name, Email, Avatar string
}

type CheckOrCreate struct {
	ID    int
	IsNew bool
}

type Authorize interface {
	CheckOrCreate(context.Context, SomeUser) (*CheckOrCreate, error)
	Authorize(int) (*ent.User, error)
}

type UserModel interface {
	Authorize
	GetUser(int) (*ent.User, error)
	GetUserAccounts(*ent.User) ([]*ent.Account, error)
	GetUserAccount(int, int) (*ent.Account, error)

	SaveUserAccount(account *ent.Account, name string, avatarID *int) (*ent.Account, error)
}

type userModelImp struct {
	userRepository          user.Repository
	accountRepository       account.Repository
	previewlyMutationClient previewly.MutationClient
}

// SaveUserAccount implements UserModel.
func (u userModelImp) SaveUserAccount(account *ent.Account, name string, avatarID *int) (*ent.Account, error) {
	return u.accountRepository.Save(account, name, avatarID)
}

// GetUserAccount implements UserModel.
func (u userModelImp) GetUserAccount(accountID int, userID int) (*ent.Account, error) {
	return u.userRepository.GetUserAccount(accountID, userID)
}

// GetUserAccounts implements UserModel.
func (u userModelImp) GetUserAccounts(user *ent.User) ([]*ent.Account, error) {
	return u.userRepository.GetUserAccounts(user)
}

func (u userModelImp) GetUser(userID int) (*ent.User, error) {
	return u.userRepository.FindByID(userID)
}

func (u userModelImp) CheckOrCreate(ctx context.Context, googleUser SomeUser) (*CheckOrCreate, error) {
	user, error := u.userRepository.FindByUUID(googleUser.UUID)
	error, notFound := error.(*ent.NotFoundError)

	if notFound {
		token, err := u.previewlyMutationClient.CreateToken(ctx)
		if err != nil {
			return nil, err
		}
		createdUser, err := u.userRepository.Create(googleUser.UUID, googleUser.Name, googleUser.Email, googleUser.Avatar)
		if err != nil {
			return nil, err
		}
		account, errorAccount := u.accountRepository.Create(createdUser.ID, googleUser.Name, googleUser.Avatar, *token)
		if errorAccount != nil {
			return nil, errorAccount
		}
		_, errorAddAccount := u.userRepository.AddAccount(createdUser, account)
		if errorAddAccount != nil {
			return nil, errorAddAccount
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
	fetchedUser, error := u.userRepository.FindByID(userID)
	if error != nil {
		return nil, error
	}
	return u.userRepository.Save(fetchedUser.Update().SetUpdatedAt(time.Now()))
}

func NewSomeUser(uuid, name, email, avatar string) SomeUser {
	return SomeUser{UUID: uuid, Name: name, Email: email, Avatar: avatar}
}

func NewUserModel(userRepository user.Repository,
	accountRepository account.Repository,
	mutationClient previewly.MutationClient,
) UserModel {
	return userModelImp{
		userRepository:          userRepository,
		accountRepository:       accountRepository,
		previewlyMutationClient: mutationClient,
	}
}
