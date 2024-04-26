package timeline

import (
	"timeline/backend/db/repository/timeline"
	"timeline/backend/ent"
)

type UserTimeline interface {
	GetUserTimelines(*ent.User) ([]*ent.Timeline, error)
	CreateTimeline(*string, *ent.User) (*ent.Timeline, error)
	GetUserTimeline(userID int, timelineID int) (*ent.Timeline, error)
	AttachEvent(*ent.Timeline, *ent.Event) (*ent.Timeline, error)
}

type timelineModelImpl struct {
	repository timeline.TimelineRepository
}

func (t timelineModelImpl) AttachEvent(timeline *ent.Timeline, event *ent.Event) (*ent.Timeline, error) {
	return t.repository.Save(timeline.Update().AddEvent(event))
}

func (t timelineModelImpl) GetUserTimeline(userID int, timelineID int) (*ent.Timeline, error) {
	return t.repository.GetUserTimeline(userID, timelineID)
}

func (t timelineModelImpl) CreateTimeline(timelineName *string, user *ent.User) (*ent.Timeline, error) {
	var name string
	if timelineName != nil {
		name = *timelineName
	} else {
		name = ""
	}
	return t.repository.Create(name, user)
}

func (t timelineModelImpl) GetUserTimelines(user *ent.User) ([]*ent.Timeline, error) {
	return t.repository.GetUserTimelines(user)
}

func NewTimelineModel(repository timeline.TimelineRepository) UserTimeline {
	return timelineModelImpl{repository: repository}
}
