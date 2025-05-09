package gettimeline

import (
	"context"

	"timeline/backend/db/model/settings"
	timelineModel "timeline/backend/db/model/timeline"
	"timeline/backend/graph/model"
	"timeline/backend/graph/resolvers"
)

type (
	resolverImpl struct {
		timelineModel timelineModel.Timeline
		settingsModel settings.Model
	}
)

// Resolve implements resolvers.Resolver.
func (r resolverImpl) Resolve(ctx context.Context, arguments resolvers.ValidArguments[ValidGetTimelineArguments]) (*model.Timeline, error) {
	args := arguments.GetArguments()
	timeline, err := r.timelineModel.GetTimeline(args.timelineID)
	if err != nil {
		return nil, err
	}

	account, err := timeline.QueryAccount().Only(ctx)
	if err != nil {
		return nil, err
	}

	return &model.Timeline{
		ID:        timeline.ID,
		Name:      &timeline.Name,
		AccountID: account.ID,
	}, nil
}

func NewResolver(timelineModel timelineModel.Timeline, settingsModel settings.Model) resolvers.Resolver[*model.Timeline, ValidGetTimelineArguments] {
	return resolverImpl{timelineModel: timelineModel, settingsModel: settingsModel}
}
