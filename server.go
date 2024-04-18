package main

import (
	"log"
	"net/http"
	"timeline/backend/config"
	"timeline/backend/db"
	"timeline/backend/graph"
	"timeline/backend/http/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()

	client := db.CreateClient(
		db.CreateConnectionURL(
			db.PostgresConfig{
				Host:     config.AppConfig.Postgres.Host,
				Port:     config.AppConfig.Postgres.Port,
				User:     config.AppConfig.Postgres.User,
				Password: config.AppConfig.Postgres.Password,
				Database: config.AppConfig.Postgres.Database,
			}))

	defer client.Close()

	router.Use(middleware.Cors(config.AppConfig.CORS.AllowedOrigin, config.AppConfig.CORS.Debug).Handler)
	router.Use(middleware.AuthMiddleware(config.AppConfig.GoogleClintID))

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/graphql", srv)

	log.Fatal(http.ListenAndServe(":"+config.AppConfig.Port, router))
}
