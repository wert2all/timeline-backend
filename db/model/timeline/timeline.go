package timeline

import (
	"timeline/backend/db/repository/timeline"
	"timeline/backend/ent"
)

type UserTimeline interface {
	GetUserTimelines(*ent.User) ([]*ent.Timeline, error)
}

type timelineModelImpl struct {
	repository timeline.TimelineRepository
}

func (t timelineModelImpl) GetUserTimelines(user *ent.User) ([]*ent.Timeline, error) {
	return t.repository.GetUserTimelines(user)
}

func NewTimelineModel(repository timeline.TimelineRepository) UserTimeline {
	return timelineModelImpl{repository: repository}
}
