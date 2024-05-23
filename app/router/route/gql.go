package route

import (
	"net/http"

	"github.com/go-chi/chi"
)

type gqlRoute struct {
	handler    http.HandlerFunc
	middleware func(http.Handler) http.Handler
}

func (g *gqlRoute) Apply(router chi.Router) {
	router.Options("/graphql", g.handler)
	router.Group(func(chiRouter chi.Router) {
		chiRouter.Use(g.middleware)
		chiRouter.Post("/graphql", g.handler)
	})
}

func NewGQLRoute(handler http.HandlerFunc, middleware func(http.Handler) http.Handler) Route {
	return &gqlRoute{handler: handler, middleware: middleware}
}
