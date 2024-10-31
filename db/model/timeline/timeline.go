package timeline

import (
	"timeline/backend/db/query"
	"timeline/backend/db/repository/timeline"
	"timeline/backend/ent"
)

type UserTimeline interface {
	GetUserTimelines(*ent.User) ([]*ent.Timeline, error)
	CreateTimeline(string, *ent.User) (*ent.Timeline, error)
	GetUserTimeline(userID int, timelineID int) (*ent.Timeline, error)
	AttachEvent(*ent.Timeline, *ent.Event) (*ent.Timeline, error)
	GetEvents(*ent.Timeline, query.Limit) ([]*ent.Event, error)
}

type timelineModelImpl struct {
	repository timeline.Repository
}

func (t timelineModelImpl) GetEvents(timeline *ent.Timeline, limit query.Limit) ([]*ent.Event, error) {
	return t.repository.GetTimelineEvents(timeline, limit)
}

func (t timelineModelImpl) AttachEvent(timeline *ent.Timeline, event *ent.Event) (*ent.Timeline, error) {
	return t.repository.Save(timeline.Update().ClearEvent().AddEvent(event))
}

func (t timelineModelImpl) GetUserTimeline(userID int, timelineID int) (*ent.Timeline, error) {
	return t.repository.GetUserTimeline(userID, timelineID)
}

func (t timelineModelImpl) CreateTimeline(timelineName string, user *ent.User) (*ent.Timeline, error) {
	return t.repository.Create(timelineName, user)
}

func (t timelineModelImpl) GetUserTimelines(user *ent.User) ([]*ent.Timeline, error) {
	return t.repository.GetUserTimelines(user)
}

func NewTimelineModel(repository timeline.Repository) UserTimeline {
	return timelineModelImpl{repository: repository}
}
