package event

import (
	"time"
	"timeline/backend/db/repository/event"
	"timeline/backend/db/repository/tag"
	"timeline/backend/ent"
	entEvent "timeline/backend/ent/event"
)

type Model interface {
	Create(date time.Time, eventType entEvent.Type) (*ent.Event, error)
	Update(*ent.EventUpdateOne) (*ent.Event, error)
	GetEventByID(id int) (*ent.Event, error)
}

type modelImpl struct {
	eventRepository event.Repository
	tagRepository   tag.Repository
}

func (m modelImpl) GetEventByID(id int) (*ent.Event, error) { return m.eventRepository.FindByID(id) }

// Update implements Model.
func (m modelImpl) Update(event *ent.EventUpdateOne) (*ent.Event, error) {
	return m.eventRepository.UpdateEvent(event)
}

func (m modelImpl) Create(date time.Time, eventType entEvent.Type) (*ent.Event, error) {
	return m.eventRepository.CreateEvent(date, eventType)
}

func NewEventModel(eventRepository event.Repository, tagRepository tag.Repository) Model {
	return modelImpl{eventRepository: eventRepository, tagRepository: tagRepository}
}
