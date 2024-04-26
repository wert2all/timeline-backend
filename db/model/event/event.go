package event

import (
	"time"
	"timeline/backend/db/repository/event"
	"timeline/backend/ent"
	entEvent "timeline/backend/ent/event"
)

type Model interface {
	Create(date time.Time, eventType entEvent.Type) (*ent.Event, error)
}

type modelImpl struct {
	repository event.Repository
}

func (m modelImpl) Create(date time.Time, eventType entEvent.Type) (*ent.Event, error) {
	return m.repository.CreateEvent(date, eventType)
}

func NewEventModel(repository event.Repository) Model {
	return modelImpl{repository: repository}
}
