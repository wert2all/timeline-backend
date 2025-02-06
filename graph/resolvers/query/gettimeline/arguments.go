package gettimeline

import "timeline/backend/graph/resolvers"

type (
	GetTimelineArgumentFactory struct{}
	GetTimelineArguments       struct {
		timelineID int
	}
)

// GetArguments implements resolvers.Arguments.
func (g GetTimelineArguments) GetArguments() GetTimelineArguments { return g }

func (g GetTimelineArgumentFactory) New(timelineID int) resolvers.Arguments[GetTimelineArguments] {
	return GetTimelineArguments{timelineID: timelineID}
}
