package settings

import (
	"context"

	"timeline/backend/ent"
	enumvalues "timeline/backend/lib/enum-values"
)

type (
	Repository interface {
		GetSettings(enumvalues.SettingsType, int) ([]*ent.Settings, error)
	}

	repositoryImpl struct {
		context context.Context
		client  *ent.Client
	}
)

// GetSettings implements Repository.
func (r repositoryImpl) GetSettings(enumvalues.SettingsType, int) ([]*ent.Settings, error) {
	panic("unimplemented")
}

func NewRepository(ctx context.Context, client *ent.Client) Repository {
	return repositoryImpl{
		context: ctx,
		client:  client,
	}
}
