package timeline

import (
	"timeline/backend/db/query"
	"timeline/backend/db/repository/timeline"
	"timeline/backend/ent"
)

type Timeline interface {
	GetAccountTimelines(*ent.Account) ([]*ent.Timeline, error)
	CreateTimeline(string, *ent.Account) (*ent.Timeline, error)
	GetTimeline(int) (*ent.Timeline, error)
	AttachEvent(*ent.Timeline, *ent.Event) (*ent.Timeline, error)
	GetEvents(*ent.Timeline, query.Limit) ([]*ent.Event, error)

	CheckUserTimeline(*ent.Timeline, int) error
}

type timelineModelImpl struct {
	repository timeline.Repository
}

// CheckUserTimeline implements Timeline.
func (t timelineModelImpl) CheckUserTimeline(timeline *ent.Timeline, userID int) error {
	return t.repository.CheckUserTimeline(timeline, userID)
}

func (t timelineModelImpl) GetEvents(timeline *ent.Timeline, limit query.Limit) ([]*ent.Event, error) {
	return t.repository.GetTimelineEvents(timeline, limit)
}

func (t timelineModelImpl) AttachEvent(timeline *ent.Timeline, event *ent.Event) (*ent.Timeline, error) {
	return t.repository.Save(timeline.Update().ClearEvent().AddEvent(event))
}

func (t timelineModelImpl) GetTimeline(timelineID int) (*ent.Timeline, error) {
	return t.repository.FindByID(timelineID)
}

func (t timelineModelImpl) CreateTimeline(timelineName string, user *ent.Account) (*ent.Timeline, error) {
	return t.repository.Create(timelineName, user)
}

func (t timelineModelImpl) GetAccountTimelines(user *ent.Account) ([]*ent.Timeline, error) {
	return t.repository.GetAccountTimelines(user)
}

func NewTimelineModel(repository timeline.Repository) Timeline {
	return timelineModelImpl{repository: repository}
}
