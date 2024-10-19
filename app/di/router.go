package di

import (
	"net/http"
	middlewares "timeline/backend/app/middleware"
	"timeline/backend/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
)

func newRouter(middlewares middlewares.Middlewares) *chi.Mux {
	gqlHandler := createGQLHandler()

	router := chi.NewRouter()
	router.Use(middlewares.List...)

	router.Options("/graphql", gqlHandler)
	router.Post("/graphql", gqlHandler)

	router.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("server panic")
	})

	return router
}

func createGQLHandler() http.HandlerFunc {
	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &graph.Resolver{}},
		),
	)
	return func(w http.ResponseWriter, r *http.Request) { srv.ServeHTTP(w, r) }
}
