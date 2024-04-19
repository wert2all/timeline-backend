package app

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"
	"timeline/backend/app/http/middleware"
	"timeline/backend/db/model/user"
	"timeline/backend/ent"
	"timeline/backend/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
)

type Application interface {
	Start()
	Stop()
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
	Client *ent.Client
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

func (a *app) Stop() {
	a.state.Client.Close()
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

func getHandlerFactory() *handlerFactory {
	return &handlerFactory{}
}

func getRouterFactory(handler http.HandlerFunc, userModel user.Authorize) *routerFactory {
	return &routerFactory{
		handler:   handler,
		userModel: userModel,
	}
}

func createConnectionURL(config Postgres) string {
	var sb strings.Builder

	optionsMap := map[string]string{
		"host":     config.Host,
		"port":     strconv.Itoa(config.Port),
		"user":     config.User,
		"password": config.Password,
		"dbname":   config.Database,
		"sslmode":  "disable",
	}

	for key, val := range optionsMap {
		sb.WriteString(key + "=" + val + " ")
	}

	return sb.String()
}
func createClient(connectionURL string) *ent.Client {
	client, err := ent.Open("postgres", connectionURL)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func NewAppState(config AppConfig) AppState {
	return AppState{Config: config, Client: createClient(createConnectionURL(config.Postgres))}
}

func NewApplication(state AppState, userModel user.Authorize) Application {
	handler := getHandlerFactory().Create(state)
	router := getRouterFactory(handler, userModel).Create(state)

	return &app{router: router, state: state}
}
