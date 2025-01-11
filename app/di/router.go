package di

import (
	"net/http"

	middlewares "timeline/backend/app/middleware"
	"timeline/backend/graph"

	"github.com/go-chi/chi/v5"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func newRouter(middlewares middlewares.Middlewares) *chi.Mux {
	srv := createGQLServer(createSchema())
	gqlHandler := func(w http.ResponseWriter, r *http.Request) { srv.ServeHTTP(w, r) }

	router := chi.NewRouter()
	router.Use(middlewares.List...)

	router.Options("/graphql", gqlHandler)
	router.Group(func(chiRouter chi.Router) {
		chiRouter.Use(middlewares.List...)
		chiRouter.Post("/graphql", gqlHandler)
	})

	router.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("server panic")
	})

	return router
}

func createSchema() graphql.ExecutableSchema {
	return graph.NewExecutableSchema(
		graph.Config{Resolvers: &graph.Resolver{}},
	)
}

func createGQLServer(schema graphql.ExecutableSchema) *handler.Server {
	srv := handler.New(schema)

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})
	return srv
}
