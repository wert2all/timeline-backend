package add

import "timeline/backend/graph/resolvers"

type (
	AddAccountArgumentFactory struct{}
	AddAccountArguments       struct {
		Name string
	}
)

// GetArguments implements resolvers.Arguments.
func (a AddAccountArguments) GetArguments() AddAccountArguments { return a }

func (a AddAccountArgumentFactory) New(name string) resolvers.Arguments[AddAccountArguments] {
	return AddAccountArguments{Name: name}
}
