package app

import (
	"log"
	"net/http"
	appHttp "timeline/backend/app/http"
	"timeline/backend/app/http/middleware"
	"timeline/backend/config"
	"timeline/backend/db/model/user"

	"github.com/go-chi/chi"
)

func Start(userModel user.UserModel) {
	handler := appHttp.Handler()

	router := chi.NewRouter()
	router.Use(middleware.Cors(config.AppConfig.CORS.AllowedOrigin, config.AppConfig.CORS.Debug).Handler)

	router.Options("/graphql", handler)
	router.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(config.AppConfig.GoogleClintID, userModel))
		r.Post("/graphql", handler)
	})

	log.Fatal(http.ListenAndServe(":"+config.AppConfig.Port, router))
}
