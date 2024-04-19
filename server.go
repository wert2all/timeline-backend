package main

import (
	"context"
	"log"
	"net/http"
	"timeline/backend/config"
	"timeline/backend/db"
	"timeline/backend/db/model/user"
	userRepository "timeline/backend/db/repository/user"
	"timeline/backend/graph"
	"timeline/backend/http/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
)

func main() {
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
	ctx := context.Background()

	userModel := user.NewUserModel(userRepository.NewUserRepository(ctx, client))

	authMiddleWare := middleware.AuthMiddleware(config.AppConfig.GoogleClintID, userModel)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	router := chi.NewRouter()
	router.Use(middleware.Cors(config.AppConfig.CORS.AllowedOrigin, config.AppConfig.CORS.Debug).Handler)
	router.Use(authMiddleWare)

	router.Post("/graphql", func(w http.ResponseWriter, r *http.Request) {
		srv.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":"+config.AppConfig.Port, router))
}
