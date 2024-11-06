package account

import (
	"context"
	"fmt"

	"timeline/backend/ent"
)

type (
	Repository interface {
		Create(int, string, string) (*ent.Account, error)
	}

	repositoryImp struct {
		context context.Context
		client  *ent.Client
	}
)

// Create implements Repository.
func (r repositoryImp) Create(userID int, name string, avatar string) (*ent.Account, error) {
	account, err := r.client.Account.Create().SetName(name).SetAvatar(avatar).Save(r.context)
	if err != nil {
		panic(fmt.Errorf("failed creating account: %w", err))
	}
	return account, nil
}

func NewRepository(ctx context.Context, client *ent.Client) Repository {
	return repositoryImp{context: ctx, client: client}
}
