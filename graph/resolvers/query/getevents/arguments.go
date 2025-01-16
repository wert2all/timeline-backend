package getevents

import (
	"timeline/backend/graph/model"
	"timeline/backend/graph/resolvers"
)

type (
	GetCursorEventsArgumentFactory struct{}
	GetCursorEventsArguments       struct {
		accountID  int
		timelineID int
		limit      *model.Limit
		cursor     *string
	}
)

func (g GetCursorEventsArgumentFactory) New(accountID int, timelineID int, limit *model.Limit, cursor *string) resolvers.Arguments[GetCursorEventsArguments] {
	return GetCursorEventsArguments{accountID: accountID, timelineID: timelineID, limit: limit, cursor: cursor}
}

func (g GetCursorEventsArguments) GetArguments() GetCursorEventsArguments { return g }
