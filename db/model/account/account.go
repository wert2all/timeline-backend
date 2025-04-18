package account

import (
	"context"

	"timeline/backend/db/repository/account"
	"timeline/backend/db/repository/user"
	"timeline/backend/ent"
	"timeline/backend/graph/client/previewly"
)

type (
	Model interface {
		NewAccount(ctx context.Context, user *ent.User, name string) (*ent.Account, error)
		GetAccount(accountID int) (*ent.Account, error)
	}
	modelImp struct {
		userRepository          user.Repository
		accountRepository       account.Repository
		previewlyMutationClient previewly.MutationClient
	}
)

// GetAccount implements Model.
func (m modelImp) GetAccount(accountID int) (*ent.Account, error) {
	return m.accountRepository.Get(accountID)
}

// New implements Model.
func (m modelImp) NewAccount(ctx context.Context, user *ent.User, name string) (*ent.Account, error) {
	token, err := m.previewlyMutationClient.CreateToken(ctx)
	if err != nil {
		return nil, err
	}
	account, errorAccount := m.accountRepository.Create(user.ID, name, *token)
	if errorAccount != nil {
		return nil, errorAccount
	}
	_, errorAddAccount := m.userRepository.AddAccount(user, account)
	if errorAddAccount != nil {
		return nil, errorAddAccount
	}
	return account, nil
}

func NewModel(accountRepository account.Repository, userRepository user.Repository, previewlyMutationClient previewly.MutationClient) Model {
	return modelImp{accountRepository: accountRepository, userRepository: userRepository, previewlyMutationClient: previewlyMutationClient}
}
