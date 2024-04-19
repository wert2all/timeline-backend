package app

import (
	"log"
	"net/http"
	"timeline/backend/app/http/middleware"
	"timeline/backend/db/model/user"
	"timeline/backend/graph"

	_ "github.com/sakirsensoy/genv/dotenv/autoload"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
	"github.com/sakirsensoy/genv"
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
	Config AppConfig
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
		r.Use(middleware.AuthMiddleware(state.Config.GoogleClintID, a.userModel))
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

func NewApplication(state AppState, userModel user.Authorize) Application {
	handler := getHandlerFactory().Create(state)
	router := getRouterFactory(handler, userModel).Create(state)

	return &app{router: router, state: state}
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

func ReadConfig() AppConfig {
	return AppConfig{
		Port: genv.Key("PORT").Default("8000").String(),
		CORS: CORS{
			Debug:         genv.Key("CORS_DEBUG").Default(false).Bool(),
			AllowedOrigin: genv.Key("CORS_ALLOWED_ORIGIN").String(),
		},
		Postgres: Postgres{
			Host:     genv.Key("POSTGRES_HOST").Default("localhost").String(),
			Port:     genv.Key("POSTGRES_PORT").Default(5432).Int(),
			User:     genv.Key("POSTGRES_USER").String(),
			Password: genv.Key("POSTGRES_PASSWORD").String(),
			Database: genv.Key("POSTGRES_DB").String(),
		},
		GoogleClintID: genv.Key("GOOGLE_CLIENT_ID").String(),
	}
}

func NewAppState(config AppConfig) AppState {
	return AppState{Config: config}
}
