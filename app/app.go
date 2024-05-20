package app

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"timeline/backend/app/http/middleware"
	"timeline/backend/db/model/user"
	"timeline/backend/di"
	"timeline/backend/graph"
)

type Application interface {
	Start()
}

type Factory[T any] interface {
	Create(state State) T
}

type State struct {
	Config di.Config
	Models di.Models
}

type app struct {
	router chi.Router
	state  State
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
func (a *routerFactory) Create(state State) chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Cors(state.Config.App.Cors.AllowedOrigin, state.Config.App.Cors.Debug).Handler)
	router.Use(middleware.Sentry())

	router.Options("/graphql", a.handler)
	router.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(a.userModel, state.Config.Google.ClientId))
		r.Post("/graphql", a.handler)
	})
	return router
}

func (h *handlerFactory) Create(state State, locator di.ServiceLocator) http.HandlerFunc {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			Models: state.Models, ServiceLocator: locator,
		},
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

func NewAppState(models di.Models, config di.Config) State {
	return State{Config: config, Models: models}
}

func NewApplication(state State, locator di.ServiceLocator) Application {
	return &app{
		router: getRouterFactory(getHandlerFactory().Create(state, locator), state.Models.Users).Create(state),
		state:  state,
	}
}
