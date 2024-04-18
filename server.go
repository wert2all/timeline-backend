package main

import (
	"log"
	"net/http"
	"timeline/backend/auth"
	"timeline/backend/config"
	"timeline/backend/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	_ "github.com/sakirsensoy/genv/dotenv/autoload"
)

func main() {
	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/graphql", srv)
	http.Handle("/graphql", srv)
	log.Fatal(http.ListenAndServe(":"+config.AppConfig.Port, router))
}
