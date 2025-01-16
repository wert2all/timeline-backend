package getevents

import (
	"context"
	"errors"

	appContext "timeline/backend/app/context"
	"timeline/backend/db/model/timeline"
	"timeline/backend/db/model/user"
	"timeline/backend/db/query"
	"timeline/backend/domain/db/cursor"
	domainUser "timeline/backend/domain/user"
	"timeline/backend/ent"
	"timeline/backend/graph/convert"
	"timeline/backend/graph/resolvers"
)

type (
	validatorImpl struct {
		userModel     user.UserModel
		timelineModel timeline.Timeline
		userExtractor domainUser.UserExtractor
	}
	ValidGetCursorEventsArguments struct {
		timeline ent.Timeline
		limit    query.Limit
		cursor   *cursor.Cursor
	}
)

func (v validatorImpl) Validate(ctx context.Context, arguments resolvers.Arguments[GetCursorEventsArguments]) (resolvers.ValidArguments[ValidGetCursorEventsArguments], error) {
	token := appContext.GetToken(ctx)
	user, errExtraction := v.userExtractor.ExtractUserFromToken(ctx, token)
	if errExtraction != nil {
		return nil, errors.New("could not expose events: " + errExtraction.Error())
	}
	args := arguments.GetArguments()
	account, err := v.userModel.GetUserAccount(args.accountID, user.ID)
	if err != nil {
		return nil, errors.New("could not expose events: " + err.Error())
	}
	timeline, errTimeline := v.timelineModel.GetAccountTimeline(account, args.timelineID)
	if errTimeline != nil {
		return nil, errors.New("could not expose events: " + errTimeline.Error())
	}
	cursor, _ := cursor.Decode(args.cursor)

	return ValidGetCursorEventsArguments{
		timeline: *timeline,
		limit:    convert.ToLimit(args.limit),
		cursor:   cursor,
	}, nil
}

func NewValidator(userModel user.UserModel, timelineModel timeline.Timeline, userExtractor domainUser.UserExtractor) resolvers.Validator[GetCursorEventsArguments, ValidGetCursorEventsArguments] {
	return validatorImpl{
		userModel:     userModel,
		timelineModel: timelineModel,
		userExtractor: userExtractor,
	}
}

func (v ValidGetCursorEventsArguments) GetArguments() ValidGetCursorEventsArguments { return v }
