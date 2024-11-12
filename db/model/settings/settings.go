package settings

import "timeline/backend/db/repository/settings"

type (
	Model     interface{}
	modelImpl struct {
		repository settings.Repository
	}
)

func NewModel(repository settings.Repository) Model {
	return modelImpl{repository: repository}
}
