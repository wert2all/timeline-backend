package app

import (
	"log"
	"net/http"
	"timeline/backend/app/http/middleware"
	"timeline/backend/db/model/user"
	"timeline/backend/di"
	"timeline/backend/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
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
	handler   http.HandlerFunc
	userModel user.Authorize
}

type handlerFactory struct{}

// Start implements Application.
func (a *app) Start() {
	log.Fatal(http.ListenAndServe(":8000", a.router))
}

// Create implements Factory.
func (a *routerFactory) Create(config di.Config) chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Cors(config.App.Cors.AllowedOrigin, config.App.Cors.Debug).Handler)
	router.Use(chiMiddleware.Recoverer)
	router.Use(middleware.Sentry())

	router.Options("/graphql", a.handler)
	router.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(a.userModel, config.Google.ClientId))
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

func getRouterFactory(handler http.HandlerFunc, userModel user.Authorize) *routerFactory {
	return &routerFactory{
		handler:   handler,
		userModel: userModel,
	}
}

func NewApplication(config di.Config, locator di.ServiceLocator) Application {
	return &app{
		router: getRouterFactory(getHandlerFactory().Create(config, locator), locator.Models().Users()).Create(config),
	}
}
