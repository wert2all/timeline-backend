package account

import (
	"context"

	"timeline/backend/ent"
)

type (
	Repository interface{}

	repositoryImp struct {
		context context.Context
		client  *ent.Client
	}
)

func NewRepository(ctx context.Context, client *ent.Client) Repository {
	return repositoryImp{context: ctx, client: client}
}
