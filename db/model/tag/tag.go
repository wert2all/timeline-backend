package tag

import (
	"timeline/backend/db/repository/tag"
	"timeline/backend/ent"
)

type Model interface {
	UpsertTag(string) (*ent.Tag, error)
}

type tagModelImpl struct {
	repository tag.Repository
}

// UpsertTag implements Model.
func (t tagModelImpl) UpsertTag(name string) (*ent.Tag, error) {
	return t.repository.Upset(name)
}

func NewTagModel(repository tag.Repository) Model {
	return tagModelImpl{repository: repository}
}
