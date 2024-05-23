package handler

import (
	"net/http"
	"timeline/backend/di"
	"timeline/backend/graph"

	"github.com/99designs/gqlgen/graphql/handler"
)

func NewGQLHandler(serviceLocator di.ServiceLocator) http.HandlerFunc {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{ServiceLocator: serviceLocator},
	}))

	return func(w http.ResponseWriter, r *http.Request) { srv.ServeHTTP(w, r) }
}
