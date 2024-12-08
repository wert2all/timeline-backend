package settings

import (
	"timeline/backend/db/repository/settings"
	"timeline/backend/ent"
	enumvalues "timeline/backend/lib/enum-values"
)

type (
	Model interface {
		GetSettings(enumvalues.SettingsType, int) []*ent.Settings
		SaveSettings(enumvalues.SettingsType, int, map[string]string) ([]*ent.Settings, error)
	}
	modelImpl struct {
		repository settings.Repository
	}
)

// SaveSettings implements Model.
func (m modelImpl) SaveSettings(entityType enumvalues.SettingsType, entityID int, settings map[string]string) ([]*ent.Settings, error) {
	return m.repository.UpsetSettings(entityType, entityID, settings)
}

// GetSettings implements Model.
func (m modelImpl) GetSettings(entityType enumvalues.SettingsType, entityID int) []*ent.Settings {
	return m.repository.GetSettings(entityType, entityID)
}

func NewModel(repository settings.Repository) Model {
	return modelImpl{repository: repository}
}
