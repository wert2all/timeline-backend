package settings

import (
	"context"

	"timeline/backend/ent"
	"timeline/backend/ent/settings"
	enumvalues "timeline/backend/lib/enum-values"
)

type (
	Repository interface {
		GetSettings(enumvalues.SettingsType, int) []*ent.Settings
	}

	repositoryImpl struct {
		context context.Context
		client  *ent.Client
	}
)

// GetSettings implements Repository.
func (r repositoryImpl) GetSettings(entityType enumvalues.SettingsType, entityID int) []*ent.Settings {
	settings, err := r.client.Settings.
		Query().
		Where(
			settings.TypeEQ(entityType),
			settings.EntityID(entityID),
		).
		All(r.context)
	if err != nil {
		return make([]*ent.Settings, 0)
	}
	return settings
}

func NewRepository(ctx context.Context, client *ent.Client) Repository {
	return repositoryImpl{
		context: ctx,
		client:  client,
	}
}
