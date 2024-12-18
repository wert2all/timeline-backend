package convert

import (
	"timeline/backend/db/query"
	"timeline/backend/ent"
	"timeline/backend/graph/model"
)

func ToEvent(event *ent.Event, tags []string) *model.TimelineEvent {
	return &model.TimelineEvent{
		ID:          event.ID,
		Date:        event.Date,
		Type:        model.TimelineType(event.Type.String()),
		Title:       &event.Title,
		Description: &event.Description,
		ShowTime:    &event.ShowTime,
		URL:         &event.URL,
		Tags:        tags,
	}
}

func ToEvents(events []*ent.Event, tags map[int][]string) []*model.TimelineEvent {
	converted := make([]*model.TimelineEvent, len(events))
	for i, event := range events {
		converted[i] = ToEvent(event, tags[event.ID])
	}
	return converted
}

func ToLimit(limit *model.Limit) query.Limit {
	if limit != nil {
		return query.NewLimit(*limit.From, *limit.To)
	}
	return query.NewLimit(0, 100)
}

func ToShortTimelines(timelines []*ent.Timeline) []*model.ShortTimeline {
	converted := make([]*model.ShortTimeline, len(timelines))
	for i, timeline := range timelines {
		converted[i] = &model.ShortTimeline{
			ID:   timeline.ID,
			Name: &timeline.Name,
		}
	}
	return converted
}
