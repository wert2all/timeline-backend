package settings

import (
	"timeline/backend/db/repository/settings"
	"timeline/backend/ent"
	enumvalues "timeline/backend/lib/enum-values"
)

type (
	Model interface {
		GetSettings(enumvalues.SettingsType, int) []*ent.Settings
	}
	modelImpl struct {
		repository settings.Repository
	}
)

// GetSettings implements Model.
func (m modelImpl) GetSettings(entityType enumvalues.SettingsType, entityID int) []*ent.Settings {
	return m.repository.GetSettings(entityType, entityID)
}

func NewModel(repository settings.Repository) Model {
	return modelImpl{repository: repository}
}
