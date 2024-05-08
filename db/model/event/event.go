package event

import (
	"time"
	"timeline/backend/db/repository/event"
	"timeline/backend/ent"
	entEvent "timeline/backend/ent/event"
)

type Model interface {
	Create(date time.Time, eventType entEvent.Type) (*ent.Event, error)
	Update(*ent.EventUpdateOne) (*ent.Event, error)
	GetEventByID(id int) (*ent.Event, error)
}

type modelImpl struct {
	repository event.Repository
}

func (m modelImpl) GetEventByID(id int) (*ent.Event, error) { return m.repository.FindByID(id) }

// Update implements Model.
func (m modelImpl) Update(event *ent.EventUpdateOne) (*ent.Event, error) {
	return m.repository.UpdateEvent(event)
}

func (m modelImpl) Create(date time.Time, eventType entEvent.Type) (*ent.Event, error) {
	return m.repository.CreateEvent(date, eventType)
}

func NewEventModel(repository event.Repository) Model { return modelImpl{repository: repository} }
