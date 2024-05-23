package app

import (
	"log"
	"net/http"
	"timeline/backend/di"
	"timeline/backend/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
)

type Application interface {
	Start()
}

type Factory[T any] interface {
	Create(di.Config) T
}

type app struct {
	router chi.Router
}

type routerFactory struct {
	handler http.HandlerFunc
}

type handlerFactory struct{}

// Start implements Application.
func (a *app) Start() {
	log.Fatal(http.ListenAndServe(":8000", a.router))
}

// Create implements Factory.
func (a *routerFactory) Create(authMiddleware func(http.Handler) http.Handler, middlewares ...func(http.Handler) http.Handler) chi.Router {
	router := chi.NewRouter()
	router.Use(middlewares...)

	router.Options("/graphql", a.handler)
	router.Group(func(r chi.Router) {
		r.Use(authMiddleware)
		r.Post("/graphql", a.handler)
	})
	return router
}

func (h *handlerFactory) Create(config di.Config, locator di.ServiceLocator) http.HandlerFunc {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{ServiceLocator: locator},
	}))

	return func(w http.ResponseWriter, r *http.Request) {
		srv.ServeHTTP(w, r)
	}
}

func getHandlerFactory() *handlerFactory {
	return &handlerFactory{}
}

func getRouterFactory(handler http.HandlerFunc) *routerFactory {
	return &routerFactory{handler: handler}
}

func NewApplication(config di.Config, locator di.ServiceLocator) Application {
	authMiddleware := locator.Middlewares().AuthMiddleware()
	middlewares := locator.Middlewares().Common()
	return &app{
		router: getRouterFactory(getHandlerFactory().Create(config, locator)).
			Create(authMiddleware, middlewares...),
	}
}
