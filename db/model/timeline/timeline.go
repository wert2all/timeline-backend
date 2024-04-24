package timeline

import (
	"timeline/backend/db/repository/timeline"
	"timeline/backend/ent"
)

type UserTimeline interface {
	GetUserTimelines(ent.User) ([]*ent.Timeline, error)
}

type TimelineModel struct {
	repository timeline.TimelineRepository
}

type timelineModelImpl struct {
	repository timeline.TimelineRepository
}

func (t timelineModelImpl) GetUserTimelines(user ent.User) ([]*ent.Timeline, error) {
	return t.GetUserTimelines(user)
}

func NewTimelineModel(repository timeline.TimelineRepository) UserTimeline {
	return timelineModelImpl{repository: repository}
}
