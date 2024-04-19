package main

import (
	"context"
	"log"
	"net/http"
	"timeline/backend/config"
	"timeline/backend/db"
	"timeline/backend/db/model/user"
	userRepository "timeline/backend/db/repository/user"
	appHttp "timeline/backend/http"
	"timeline/backend/http/middleware"

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

	router := chi.NewRouter()
	router.Use(middleware.Cors(config.AppConfig.CORS.AllowedOrigin, config.AppConfig.CORS.Debug).Handler)

	handler := appHttp.Handler()

	router.Options("/graphql", handler)
	router.Group(func(r chi.Router) {
		r.Use(authMiddleWare)
		r.Post("/graphql", handler)
	})

	log.Fatal(http.ListenAndServe(":"+config.AppConfig.Port, router))
}
