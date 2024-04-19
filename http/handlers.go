package http

import (
	"net/http"
	"timeline/backend/graph"

	"github.com/99designs/gqlgen/graphql/handler"
)

func Handler() http.HandlerFunc {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(w http.ResponseWriter, r *http.Request) {
		srv.ServeHTTP(w, r)
	}
}
