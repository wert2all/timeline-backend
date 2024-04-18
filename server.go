package main

import (
	"log"
	"net/http"
	"timeline/backend/config"
	"timeline/backend/graph"
	"timeline/backend/http/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
	_ "github.com/sakirsensoy/genv/dotenv/autoload"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Cors(config.AppConfig.CORS.AllowedOrigin, config.AppConfig.CORS.Debug).Handler)
	router.Use(middleware.AuthMiddleware())

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/graphql", srv)
	http.Handle("/graphql", srv)
	log.Fatal(http.ListenAndServe(":"+config.AppConfig.Port, router))
}
