package router

import (
	"net/http"
	"timeline/backend/app/router/route"

	"github.com/go-chi/chi/v5"
)

type RouterBuilder interface {
	SetMiddlewares(middlewares ...func(http.Handler) http.Handler) RouterBuilder
	SetRoutes(routes ...route.Route) RouterBuilder
	Build() chi.Router
}

type routerBuilder struct {
	middlewares []func(http.Handler) http.Handler
	routes      []route.Route
}

// SetRoutes implements RouterBuilder.
func (r *routerBuilder) SetRoutes(routes ...route.Route) RouterBuilder {
	r.routes = append(r.routes, routes...)
	return r
}

// SetMiddlewares implements RouterBuilder.
func (r *routerBuilder) SetMiddlewares(middlewares ...func(http.Handler) http.Handler) RouterBuilder {
	r.middlewares = append(r.middlewares, middlewares...)
	return r
}

func NewRouterBuilder() RouterBuilder {
	return &routerBuilder{
		middlewares: []func(http.Handler) http.Handler{},
		routes:      []route.Route{},
	}
}

func (r *routerBuilder) Build() chi.Router {
	router := chi.NewRouter()
	router.Use(r.middlewares...)
	for _, route := range r.routes {
		route.Apply(router)
	}
	return router
}
