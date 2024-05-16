package tag

import (
	"context"
	"timeline/backend/ent"
	"timeline/backend/ent/event"
	"timeline/backend/ent/tag"

	"entgo.io/ent/dialect/sql"
)

type Repository interface {
	FindByID(ID int) (*ent.Tag, error)
	Upset(string) (*ent.Tag, error)
	GetEventTags(int) []*ent.Tag
}

type repositoryImpl struct {
	context context.Context
	client  *ent.Client
}

func (r repositoryImpl) GetEventTags(eventID int) []*ent.Tag {
	tags, err := r.client.Tag.Query().Where(tag.HasEventWith(event.ID(eventID))).All(r.context)
	if err != nil {
		return make([]*ent.Tag, 0)
	}
	return tags
}

func (r repositoryImpl) FindByID(ID int) (*ent.Tag, error) {
	return r.client.Tag.Query().Where(tag.ID(ID)).Only(r.context)
}

func (r repositoryImpl) Upset(name string) (*ent.Tag, error) {
	entityID, err := r.client.Tag.Create().
		SetTag(name).
		OnConflict(sql.ConflictColumns(tag.FieldTag)).
		UpdateNewValues().
		ID(r.context)
	if err != nil {
		return nil, err
	}

	return r.FindByID(entityID)
}

func NewRepository(ctx context.Context, client *ent.Client) Repository {
	return repositoryImpl{context: ctx, client: client}
}
