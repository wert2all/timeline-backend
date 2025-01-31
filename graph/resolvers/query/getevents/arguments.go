package getevents

import (
	"timeline/backend/graph/model"
	"timeline/backend/graph/resolvers"
)

type (
	GetCursorEventsArgumentFactory struct{}
	GetCursorEventsArguments       struct {
		timelineID int
		accountID  *int
		limit      *model.Limit
		cursor     *string
	}
)

func (g GetCursorEventsArgumentFactory) New(timelineID int, accountID *int, limit *model.Limit, cursor *string) resolvers.Arguments[GetCursorEventsArguments] {
	return GetCursorEventsArguments{timelineID: timelineID, accountID: accountID, limit: limit, cursor: cursor}
}

func (g GetCursorEventsArguments) GetArguments() GetCursorEventsArguments { return g }
