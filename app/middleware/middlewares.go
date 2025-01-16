package middlewares

import (
	"encoding/base64"
	"net/http"
	"strings"

	appContext "timeline/backend/app/context"

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
			saveTokenMiddleware(),
		},
	}
}

func saveTokenMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			token := extractToken(req)
			contextWithData := appContext.SetToken(req.Context(), token)
			next.ServeHTTP(w, req.WithContext(contextWithData))
		})
	}
}

func extractToken(req *http.Request) *string {
	authHeader := req.Header.Get("Authorization")
	splitted := strings.Split(authHeader, "Bearer ")

	if len(splitted) == 2 {
		decoded, _ := base64.StdEncoding.DecodeString(splitted[1])
		result := string(decoded)
		return &result
	} else {
		return nil
	}
}
