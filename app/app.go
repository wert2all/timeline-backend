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

type CORS struct {
	AllowedOrigin string
	Debug         bool
}

type Postgres struct {
	Host, User, Password, Database string
	Port                           int
}

type Config struct {
	Port          string
	CORS          CORS
	Postgres      Postgres
	GoogleClintID string
	SentryDsn     string
}

type State struct {
	Config    Config
	Models    di.Models
	Resolvers di.Resolvers
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
	log.Fatal(http.ListenAndServe(":"+a.state.Config.Port, a.router))
}

// Create implements Factory.
func (a *routerFactory) Create(state State) chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Cors(state.Config.CORS.AllowedOrigin, state.Config.CORS.Debug).Handler)
	router.Use(middleware.Sentry())

	router.Options("/graphql", a.handler)
	router.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(a.userModel, state.Config.GoogleClintID))
		r.Post("/graphql", a.handler)
	})
	return router
}

func (h *handlerFactory) Create(state State) http.HandlerFunc {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			Models: state.Models, Resolvers: state.Resolvers,
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

func NewAppState(models di.Models, config Config, resolvers di.Resolvers) State {
	return State{
		Config:    config,
		Models:    models,
		Resolvers: resolvers,
	}
}

func NewApplication(state State) Application {
	return &app{
		router: getRouterFactory(getHandlerFactory().Create(state), state.Models.Users).Create(state),
		state:  state,
	}
}
