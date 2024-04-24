package timeline

import (
	"context"
	"timeline/backend/ent"
)

type TimelineRepository interface {
	GetUserTimelines(*ent.User) ([]*ent.Timeline, error)
	Create(name string, user *ent.User) (*ent.Timeline, error)
}

type timelineRepositoryImpl struct {
	client  *ent.Client
	context context.Context
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
