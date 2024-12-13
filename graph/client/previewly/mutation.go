package previewly

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

type (
	MutationClient interface {
		CreateToken(context.Context) (*string, error)
	}
	mutationImpl struct {
		client graphql.Client
	}
)

// CreateToken implements MutationClient.
func (m mutationImpl) CreateToken(ctx context.Context) (*string, error) {
	responce, err := createToken(ctx, m.client)
	if err != nil {
		return nil, err
	} else {
		return &responce.CreateToken, nil
	}
}

func NewMutationClient(client graphql.Client) MutationClient {
	return mutationImpl{client: client}
}
