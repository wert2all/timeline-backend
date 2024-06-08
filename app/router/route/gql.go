package route

import (
	"net/http"
	"timeline/backend/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
)

type gqlRoute struct {
	config     graph.Config
	middleware func(http.Handler) http.Handler
}

func (g *gqlRoute) Apply(router chi.Router) {
	handler := g.createHandler(g.config)

	router.Options("/graphql", handler)
	router.Group(func(chiRouter chi.Router) {
		chiRouter.Use(g.middleware)
		chiRouter.Post("/graphql", handler)
	})
}

func NewGQLRoute(config graph.Config, middleware func(http.Handler) http.Handler) Route {
	return &gqlRoute{middleware: middleware, config: config}
}

func NewGQLConfig(resolver *graph.Resolver) graph.Config { return graph.Config{Resolvers: resolver} }

func (g *gqlRoute) createHandler(config graph.Config) http.HandlerFunc {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(config))
	return func(w http.ResponseWriter, r *http.Request) { srv.ServeHTTP(w, r) }
}
