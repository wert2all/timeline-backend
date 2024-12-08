package settings

import (
	"context"

	"timeline/backend/ent"
	"timeline/backend/ent/settings"
	enumvalues "timeline/backend/lib/enum-values"

	"entgo.io/ent/dialect/sql"
)

type (
	Repository interface {
		GetSettings(enumvalues.SettingsType, int) []*ent.Settings
		UpsetSettings(enumvalues.SettingsType, int, map[string]string) error
	}

	repositoryImpl struct {
		context context.Context
		client  *ent.Client
	}
)

// UpsetSettings implements Repository.
func (r repositoryImpl) UpsetSettings(entityType enumvalues.SettingsType, entityID int, settingsValues map[string]string) error {
	settingsBuilders := make([]*ent.SettingsCreate, len(settingsValues))

	i := 0
	for key, value := range settingsValues {
		settingsBuilders[i] = r.client.Settings.Create().
			SetType(entityType).
			SetEntityID(entityID).
			SetKey(key).
			SetValue(value)
		i++
	}
	return r.client.Settings.
		CreateBulk(settingsBuilders...).
		OnConflict(
			sql.ConflictColumns(settings.FieldType, settings.FieldEntityID, settings.FieldKey),
		).
		UpdateValue().
		Exec(r.context)
}

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
	return repositoryImpl{context: ctx, client: client}
}
