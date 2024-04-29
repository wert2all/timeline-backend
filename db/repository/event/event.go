package event

import (
	"context"
	"time"
	"timeline/backend/ent"
	"timeline/backend/ent/event"
)

type Repository interface {
	FindByID(ID int) (*ent.Event, error)
	CreateEvent(date time.Time, eventType event.Type) (*ent.Event, error)
	UpdateEvent(*ent.EventUpdateOne) (*ent.Event, error)
}

type repositoryImpl struct {
	client  *ent.Client
	context context.Context
}

func (r repositoryImpl) UpdateEvent(event *ent.EventUpdateOne) (*ent.Event, error) {
	return event.Save(r.context)
}

func (r repositoryImpl) CreateEvent(date time.Time, eventType event.Type) (*ent.Event, error) {
	return r.client.Event.Create().SetDate(date).SetType(eventType).Save(r.context)
}

func (r repositoryImpl) FindByID(ID int) (*ent.Event, error) {
	return r.client.Event.Query().Where(event.ID(ID)).Only(r.context)
}

func NewRepository(ctx context.Context, client *ent.Client) Repository {
	return repositoryImpl{context: ctx, client: client}
}
