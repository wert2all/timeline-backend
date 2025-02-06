package gettimeline

import (
	"context"

	"timeline/backend/graph/resolvers"
)

type (
	ValidGetTimelineArguments struct{ timelineID int }
	validatorImpl             struct{}
)

// GetArguments implements resolvers.ValidArguments.
func (v ValidGetTimelineArguments) GetArguments() ValidGetTimelineArguments { return v }

// Validate implements resolvers.Validator.
func (v validatorImpl) Validate(ctx context.Context, arguments resolvers.Arguments[GetTimelineArguments]) (resolvers.ValidArguments[ValidGetTimelineArguments], error) {
	return ValidGetTimelineArguments{timelineID: arguments.GetArguments().timelineID}, nil
}

func NewValidator() resolvers.Validator[GetTimelineArguments, ValidGetTimelineArguments] {
	return validatorImpl{}
}
