package graph

import (
	"timeline/backend/db/model"
	"timeline/backend/graph/resolvers"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Models    model.AppModels
	Resolvers resolvers.Resolvers
}
