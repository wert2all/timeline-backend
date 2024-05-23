package router

import (
	"net/http"
	"timeline/backend/app/router/handler"
	"timeline/backend/di"

	"github.com/go-chi/chi"
)

type RouterFactory interface {
	Create() chi.Router
}

type routerFactory struct {
	locator          di.ServiceLocator
	authMiddleware   func(http.Handler) http.Handler
	commonMiddlwares []func(http.Handler) http.Handler
}

func NewRouterFactory(locator di.ServiceLocator) RouterFactory {
	return &routerFactory{
		locator:          locator,
		authMiddleware:   locator.Middlewares().AuthMiddleware(),
		commonMiddlwares: locator.Middlewares().Common(),
	}
}

func (r *routerFactory) Create() chi.Router {
	router := chi.NewRouter()
	router.Use(r.commonMiddlwares...)

	return r.addGQLRoute(router)
}

// TODO: refactor to abstract
func (r *routerFactory) addGQLRoute(router chi.Router) chi.Router {
	handler := handler.NewGQLHandler(r.locator)

	router.Options("/graphql", handler)
	router.Group(func(chiRouter chi.Router) {
		chiRouter.Use(r.authMiddleware)
		chiRouter.Post("/graphql", handler)
	})
	return router
}
