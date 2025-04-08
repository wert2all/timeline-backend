package getevent

import "timeline/backend/graph/resolvers"

type (
	GetEventArgumentFactory struct{}
	GetEventArguments       struct {
		eventID   int
		accountID *int
	}
)

func (g GetEventArguments) GetArguments() GetEventArguments { return g }

func (g GetEventArgumentFactory) New(eventID int, accountID *int) resolvers.Arguments[GetEventArguments] {
	return GetEventArguments{eventID: eventID, accountID: accountID}
}
