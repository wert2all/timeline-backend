package getevents

import (
	"context"
	"errors"

	appContext "timeline/backend/app/context"
	"timeline/backend/db/model/timeline"
	"timeline/backend/db/model/user"
	"timeline/backend/domain/db/cursor"
	"timeline/backend/domain/gql/validator"
	domainUser "timeline/backend/domain/user"
	"timeline/backend/ent"
	"timeline/backend/graph/resolvers"
)

type (
	validatorImpl struct {
		userModel     user.UserModel
		timelineModel timeline.Timeline
		userExtractor domainUser.UserExtractor
	}
	ValidGetCursorEventsArguments struct {
		timeline    ent.Timeline
		withPrivate bool
		cursor      *cursor.Cursor
		limit       int
	}
)

func (v validatorImpl) Validate(ctx context.Context, arguments resolvers.Arguments[GetCursorEventsArguments]) (resolvers.ValidArguments[ValidGetCursorEventsArguments], error) {
	args := arguments.GetArguments()
	timeline, errTimeline := v.timelineModel.GetTimeline(args.timelineID)
	if errTimeline != nil {
		return nil, errors.New("could not expose events: " + errTimeline.Error())
	}
	cursor, _ := cursor.Decode(args.cursor)

	return ValidGetCursorEventsArguments{
		timeline:    *timeline,
		withPrivate: v.withPrivate(ctx, v.extractAccountID(ctx, args.accountID), *timeline),
		cursor:      cursor,
		limit:       validator.NewLimit(args.limit),
	}, nil
}

func (v validatorImpl) withPrivate(ctx context.Context, accountID *int, timeline ent.Timeline) bool {
	if accountID != nil {
		timelineAccountID, err := timeline.QueryAccount().OnlyID(ctx)
		if err != nil {
			return false
		}
		return *accountID == timelineAccountID
	}
	return false
}

func (v validatorImpl) extractAccountID(ctx context.Context, requestAccountID *int) *int {
	token := appContext.GetToken(ctx)
	user, errExtraction := v.userExtractor.ExtractUserFromToken(ctx, token)
	if errExtraction != nil {
		return nil
	}
	if requestAccountID != nil {
		account, err := v.userModel.GetUserAccount(*requestAccountID, user.ID)
		if err != nil {
			return nil
		}
		return &account.ID
	}
	return nil
}

func NewValidator(userModel user.UserModel, timelineModel timeline.Timeline, userExtractor domainUser.UserExtractor) resolvers.Validator[GetCursorEventsArguments, ValidGetCursorEventsArguments] {
	return validatorImpl{
		userModel:     userModel,
		timelineModel: timelineModel,
		userExtractor: userExtractor,
	}
}

func (v ValidGetCursorEventsArguments) GetArguments() ValidGetCursorEventsArguments { return v }
