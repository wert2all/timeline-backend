package app

import (
	"log"
	"net/http"
	appContext "timeline/backend/app/context"
	"timeline/backend/app/http/middleware"
	"timeline/backend/db/model/user"
	"timeline/backend/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
)

type Application interface {
	Start()
}

type Factory[T any] interface {
	Create(state AppState) T
}

type CORS struct {
	AllowedOrigin string
	Debug         bool
}

type Postgres struct {
	Host, User, Password, Database string
	Port                           int
}

type AppConfig struct {
	Port          string
	CORS          CORS
	Postgres      Postgres
	GoogleClintID string
}

type AppState struct {
	Config  AppConfig
	Context appContext.AppContext
}

type app struct {
	router chi.Router
	state  AppState
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

// Start implements Factory.
func (a *routerFactory) Create(state AppState) chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Cors(state.Config.CORS.AllowedOrigin, state.Config.CORS.Debug).Handler)

	router.Options("/graphql", a.handler)
	router.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(state.Context, state.Config.GoogleClintID))
		r.Post("/graphql", a.handler)
	})
	return router
}

func (h *handlerFactory) Create(state AppState) http.HandlerFunc {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

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

func NewAppState(context appContext.AppContext, config AppConfig) AppState {
	return AppState{
		Config:  config,
		Context: context,
	}
}

func NewApplication(state AppState) Application {
	handler := getHandlerFactory().Create(state)
	router := getRouterFactory(handler, state.Context.GetModels().Users).Create(state)

	return &app{router: router, state: state}
}
