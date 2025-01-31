package timeline

import (
	"context"

	"timeline/backend/ent"
	"timeline/backend/ent/timeline"
	"timeline/backend/ent/user"
)

type Repository interface {
	FindByID(ID int) (*ent.Timeline, error)

	GetAccountTimelines(*ent.Account) ([]*ent.Timeline, error)
	GetAccountTimeline(account *ent.Account, timelineID int) (*ent.Timeline, error)

	CheckUserTimeline(*ent.Timeline, int) error

	Create(name string, user *ent.Account) (*ent.Timeline, error)
	Save(*ent.TimelineUpdateOne) (*ent.Timeline, error)
}

type timelineRepositoryImpl struct {
	client  *ent.Client
	context context.Context
}

// CheckUserTimeline implements Repository.
func (t timelineRepositoryImpl) CheckUserTimeline(timeline *ent.Timeline, userID int) error {
	_, err := timeline.QueryAccount().QueryUser().Where(user.ID(userID)).Only(t.context)
	return err
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

func (t timelineRepositoryImpl) Create(name string, account *ent.Account) (*ent.Timeline, error) {
	return t.client.Timeline.Create().SetName(name).SetAccount(account).Save(t.context)
}

func (t timelineRepositoryImpl) GetAccountTimelines(user *ent.Account) ([]*ent.Timeline, error) {
	return t.client.Account.QueryTimeline(user).All(t.context)
}

func (t timelineRepositoryImpl) GetAccountTimeline(account *ent.Account, timelineID int) (*ent.Timeline, error) {
	return t.client.Account.QueryTimeline(account).Where(timeline.ID(timelineID)).Only(t.context)
}

func NewTimelineRepository(ctx context.Context, client *ent.Client) Repository {
	return timelineRepositoryImpl{client: client, context: ctx}
}
