package timeline

import (
	"context"

	"timeline/backend/db/model/timeline"
	"timeline/backend/db/model/user"
	"timeline/backend/graph/convert"
	"timeline/backend/graph/model"
)

type (
	Resolver interface {
		Resolve(context.Context, int, int) ([]*model.ShortTimeline, error)
	}

	resolverImpl struct {
		timelineModel timeline.Timeline
		userModel     user.UserModel
	}
)

// Resolve implements Resolver.
func (r resolverImpl) Resolve(ctx context.Context, accountID int, userID int) ([]*model.ShortTimeline, error) {
	account, err := r.userModel.GetUserAccount(accountID, userID)
	if err != nil {
		return nil, err
	}
	timelines, errTimelines := r.timelineModel.GetAccountTimelines(account)
	if errTimelines != nil {
		return nil, errTimelines
	}

	return convert.ToShortTimelines(timelines), nil
}

func NewMyAccountTimelinesResolver(timelineModel timeline.Timeline, userModel user.UserModel) Resolver {
	return resolverImpl{
		timelineModel: timelineModel,
		userModel:     userModel,
	}
}
