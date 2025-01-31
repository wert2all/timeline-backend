package event

import (
	"time"

	"timeline/backend/db/model/tag"
	"timeline/backend/db/model/timeline"
	"timeline/backend/db/repository/event"
	"timeline/backend/domain/db/cursor"
	"timeline/backend/ent"
	entEvent "timeline/backend/ent/event"
	eventValidator "timeline/backend/graph/resolvers/mutation/event"
)

type Model interface {
	Create(date time.Time, eventType entEvent.Type) (*ent.Event, error)
	Update(*ent.EventUpdateOne) (*ent.Event, error)
	GetEventByID(int) (*ent.Event, error)

	GetTimelineEvents(timeline ent.Timeline, withPrivate bool, cursor *cursor.Cursor, limit int) ([]*ent.Event, error)

	UpdateEvent(*ent.Event, *eventValidator.BaseValidEventInput) (*ent.Event, error)
}

type modelImpl struct {
	eventRepository event.Repository

	tagModel      tag.Model
	timelineModel timeline.Timeline
}

// GetTimelineEvents implements Model.
func (m modelImpl) GetTimelineEvents(timeline ent.Timeline, withPrivate bool, cursor *cursor.Cursor, limit int) ([]*ent.Event, error) {
	return m.eventRepository.FindTimelineEventsByCursor(timeline.ID, withPrivate, cursor, limit)
}

func (m modelImpl) UpdateEvent(event *ent.Event, input *eventValidator.BaseValidEventInput) (*ent.Event, error) {
	tags := make([]*ent.Tag, 0)
	for _, tagArgument := range input.Tags {
		tagEntity, err := m.tagModel.UpsertTag(tagArgument)
		if err == nil {
			tags = append(tags, tagEntity)
		}
	}
	var shouldUpdateEntity *ent.EventUpdateOne
	shouldUpdateEntity = event.Update().
		SetTitle(input.Title).
		SetDescription(input.Description).
		SetShowTime(input.ShowTime).
		SetDate(input.Date).
		AddTags(tags...).
		ClearTimeline()

	if input.Url != nil {
		shouldUpdateEntity = shouldUpdateEntity.SetURL(input.Url.String())
	}
	if input.PreviewlyImageID != nil {
		shouldUpdateEntity = shouldUpdateEntity.SetPreviewlyImageID(*input.PreviewlyImageID)
	} else {
		shouldUpdateEntity = shouldUpdateEntity.ClearPreviewlyImageID()
	}
	updatedEntity, updateErr := m.Update(shouldUpdateEntity)

	if updateErr != nil {
		return nil, updateErr
	}

	_, err := m.timelineModel.AttachEvent(input.Timeline, updatedEntity)
	if err != nil {
		return nil, err
	}

	return updatedEntity, nil
}

func (m modelImpl) GetEventByID(id int) (*ent.Event, error) { return m.eventRepository.FindByID(id) }

func (m modelImpl) Update(event *ent.EventUpdateOne) (*ent.Event, error) {
	return m.eventRepository.UpdateEvent(event)
}

func (m modelImpl) Create(date time.Time, eventType entEvent.Type) (*ent.Event, error) {
	return m.eventRepository.CreateEvent(date, eventType)
}

func NewEventModel(eventRepository event.Repository, tagModel tag.Model, timelineModel timeline.Timeline) Model {
	return modelImpl{eventRepository: eventRepository, tagModel: tagModel, timelineModel: timelineModel}
}
