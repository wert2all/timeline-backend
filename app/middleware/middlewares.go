package middlewares

import (
	"net/http"
	"strings"

	appContext "timeline/backend/app/context"

	domainUser "timeline/backend/domain/user"

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

func NewAuthMiddleware(domainUserExtractor domainUser.UserExtractor) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// fmt.Printf("User ID: %d\n", context.GetUserID())
			token := extractToken(req)
			user, err := domainUserExtractor.ExtractUserFromToken(req.Context(), &token)
			if err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}
			contextWithdata := appContext.SetUserID(req.Context(), user.ID, user.IsNew, token)
			next.ServeHTTP(w, req.WithContext(contextWithdata))
		})
	}
}

func extractToken(req *http.Request) string {
	authHeader := req.Header.Get("Authorization")
	splitted := strings.Split(authHeader, "Bearer ")

	if len(splitted) == 2 {
		return splitted[1]
	} else {
		return ""
	}
}
