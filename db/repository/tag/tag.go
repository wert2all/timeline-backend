package tag

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"timeline/backend/ent"
	"timeline/backend/ent/tag"
)

type Repository interface {
	FindByID(ID int) (*ent.Tag, error)
	Upset(string, context.Context, ) (*ent.Tag, error)
}

type repositoryImpl struct {
	context context.Context
	client  *ent.Client
}

func (r repositoryImpl) FindByID(ID int) (*ent.Tag, error) {
	return r.client.Tag.Query().Where(tag.ID(ID)).Only(r.context)
}

func (r repositoryImpl) Upset(name string, ctx context.Context) (*ent.Tag, error) {
	entityId, err := r.client.Tag.Create().SetTag(name).OnConflict(sql.ConflictColumns(tag.FieldTag)).UpdateNewValues().ID(ctx)
	if err != nil {
		return nil, err
	}
	return r.FindByID(entityId)
}

func NewRepository(ctx context.Context, client *ent.Client) Repository {
	return repositoryImpl{context: ctx, client: client}
}
