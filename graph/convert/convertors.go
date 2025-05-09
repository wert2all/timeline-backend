package convert

import (
	"timeline/backend/db/query"
	"timeline/backend/ent"
	"timeline/backend/graph/model"
)

func ToEvent(event *ent.Event, tags []string, timelineID int) *model.TimelineEvent {
	return &model.TimelineEvent{
		ID:               event.ID,
		Date:             event.Date,
		Type:             model.TimelineType(event.Type.String()),
		Title:            &event.Title,
		TimelineID:       timelineID,
		Description:      &event.Description,
		ShowTime:         &event.ShowTime,
		URL:              &event.URL,
		Tags:             tags,
		PreviewlyImageID: event.PreviewlyImageID,
	}
}

func ToEvents(events []*ent.Event, tags map[int][]string, timelineID int) []*model.TimelineEvent {
	converted := make([]*model.TimelineEvent, len(events))
	for i, event := range events {
		converted[i] = ToEvent(event, tags[event.ID], timelineID)
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

func ToShortAccount(accountEntity ent.Account, settings []*ent.Settings) *model.ShortAccount {
	gqlSettings := make([]*model.AccountSettings, len(settings))

	for i, setting := range settings {
		gqlSettings[i] = &model.AccountSettings{
			Key:   setting.Key,
			Value: setting.Value,
		}
	}

	return &model.ShortAccount{
		Name:           &accountEntity.Name,
		ID:             accountEntity.ID,
		PreviewlyToken: accountEntity.PreviewlyToken,
		AvatarID:       accountEntity.AvatarID,
		About:          accountEntity.About,
		Settings:       gqlSettings,
	}
}
