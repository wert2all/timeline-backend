package convert

import (
	"timeline/backend/db/query"
	"timeline/backend/ent"
	"timeline/backend/graph/model"
)

func ToEvent(event *ent.Event) *model.TimelineEvent {
	return &model.TimelineEvent{
		ID:          event.ID,
		Date:        event.Date,
		Type:        model.TimelineType(event.Type.String()),
		Title:       &event.Title,
		Description: &event.Description,
		ShowTime:    &event.ShowTime,
		URL:         &event.URL,
	}
}

func ToEvents(events []*ent.Event) []*model.TimelineEvent {
	converted := make([]*model.TimelineEvent, len(events))
	for i, event := range events {
		converted[i] = ToEvent(event)
	}
	return converted
}

func ToLimit(limit *model.Limit) query.Limit {
	if limit != nil {
		return query.NewLimit(*limit.From, *limit.To)
	}
	return query.NewLimit(0, 100)
}
