package convert

import (
	"timeline/backend/ent"
	"timeline/backend/graph/model"
)

func ToUser(user *ent.User, timelines []*ent.Timeline, isNew bool) *model.User {
	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Avatar:    user.Avatar,
		IsNew:     isNew,
		Timelines: converTimelines(timelines),
	}
}

func ToShortTimeline(timeline *ent.Timeline) *model.ShortUserTimeline {
	return &model.ShortUserTimeline{ID: timeline.ID, Name: &timeline.Name}
}

func converTimelines(timelines []*ent.Timeline) []*model.ShortUserTimeline {
	gqlTimelines := make([]*model.ShortUserTimeline, len(timelines))
	for i, timeline := range timelines {
		gqlTimelines[i] = ToShortTimeline(timeline)
	}
	return gqlTimelines
}

func ToEvent(event *ent.Event) *model.TimelineEvent {
	return &model.TimelineEvent{
		ID:          event.ID,
		Date:        event.Date,
		Type:        model.TimelineType(event.Type.String()),
		Title:       &event.Title,
		Description: &event.Description,
	}
}
