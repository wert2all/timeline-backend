package middlewares

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

type (
	Middlewares struct {
		List []func(http.Handler) http.Handler
	}
)

func NewMiddlewares() Middlewares {
	return Middlewares{
		List: []func(http.Handler) http.Handler{
			middleware.Logger,
			middleware.Recoverer,
			middleware.RealIP,
			cors.New(cors.Options{
				AllowedOrigins:     []string{"*"},
				AllowCredentials:   true,
				AllowedMethods:     []string{"GET", "POST", "OPTIONS"},
				AllowedHeaders:     []string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept", "Authorization"},
				OptionsPassthrough: true,
				Debug:              true,
			}).
				Handler,
		},
	}
}
