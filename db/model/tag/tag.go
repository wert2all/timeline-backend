package tag

import (
	"timeline/backend/db/repository/tag"
	"timeline/backend/ent"
)

type Model interface {
	UpsertTag(string) (*ent.Tag, error)
	GetEventTags(*ent.Event) []*ent.Tag
}

type tagModelImpl struct {
	repository tag.Repository
}

func (t tagModelImpl) GetEventTags(event *ent.Event) []*ent.Tag {
	return t.repository.GetEventTags(event.ID)
}

// UpsertTag implements Model.
func (t tagModelImpl) UpsertTag(name string) (*ent.Tag, error) {
	return t.repository.Upset(name)
}

func NewTagModel(repository tag.Repository) Model {
	return tagModelImpl{repository: repository}
}
