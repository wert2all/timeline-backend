package account

import (
	"context"
	"fmt"

	"timeline/backend/ent"
)

type (
	Repository interface {
		Create(int, string, string, string) (*ent.Account, error)

		Save(account *ent.Account, name string, avatarID *int) (*ent.Account, error)
	}

	repositoryImp struct {
		context context.Context
		client  *ent.Client
	}
)

// Save implements Repository.
func (r repositoryImp) Save(account *ent.Account, name string, avatarID *int) (*ent.Account, error) {
	return r.client.Account.UpdateOne(account).SetName(name).Save(r.context)
}

// Create implements Repository.
func (r repositoryImp) Create(userID int,
	name string,
	avatar string,
	previewlyToken string,
) (*ent.Account, error) {
	account, err := r.client.Account.Create().
		SetName(name).
		SetAvatar(avatar).
		SetPreviewlyToken(previewlyToken).
		Save(r.context)
	if err != nil {
		panic(fmt.Errorf("failed creating account: %w", err))
	}
	return account, nil
}

func NewRepository(ctx context.Context, client *ent.Client) Repository {
	return repositoryImp{context: ctx, client: client}
}
