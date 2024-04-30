package timeline

import (
	"context"
	"timeline/backend/db/query"
	"timeline/backend/ent"
	"timeline/backend/ent/timeline"
	"timeline/backend/ent/user"
)

type TimelineRepository interface {
	FindByID(ID int) (*ent.Timeline, error)

	GetUserTimelines(*ent.User) ([]*ent.Timeline, error)
	GetUserTimeline(userID int, timelineID int) (*ent.Timeline, error)
	GetTimelineEvents(*ent.Timeline, query.Limit) ([]*ent.Event, error)

	Create(name string, user *ent.User) (*ent.Timeline, error)
	Save(*ent.TimelineUpdateOne) (*ent.Timeline, error)
}

type timelineRepositoryImpl struct {
	client  *ent.Client
	context context.Context
}

func (t timelineRepositoryImpl) GetTimelineEvents(timeline *ent.Timeline, limit query.Limit) ([]*ent.Event, error) {
	return t.client.Timeline.QueryEvent(timeline).Offset(limit.Offset).Limit(limit.Limit).All(t.context)
}

func (t timelineRepositoryImpl) GetUserTimeline(userID int, timelineID int) (*ent.Timeline, error) {
	return t.client.Timeline.Query().Where(
		timeline.And(
			timeline.ID(timelineID),
			timeline.HasUserWith(user.ID(userID)),
		)).Only(t.context)
}

func (t timelineRepositoryImpl) Save(timeline *ent.TimelineUpdateOne) (*ent.Timeline, error) {
	return timeline.Save(t.context)
}

func (t timelineRepositoryImpl) FindByID(ID int) (*ent.Timeline, error) {
	return t.client.Timeline.Query().Where(timeline.ID(ID)).Only(t.context)
}

func (t timelineRepositoryImpl) GetTimeline(timelineID int) (*ent.Timeline, error) {
	return t.client.Timeline.Query().Where(timeline.ID(timelineID)).Only(t.context)
}

func (t timelineRepositoryImpl) Create(name string, user *ent.User) (*ent.Timeline, error) {
	return t.client.Timeline.Create().SetName(name).SetUser(user).Save(t.context)

}

func (t timelineRepositoryImpl) GetUserTimelines(user *ent.User) ([]*ent.Timeline, error) {
	return t.client.User.QueryTimeline(user).All(t.context)
}

func NewTimelineRepository(ctx context.Context, client *ent.Client) TimelineRepository {
	return timelineRepositoryImpl{client: client, context: ctx}
}
