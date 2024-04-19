package middleware

import (
	"context"
	"net/http"
	"strings"
	appContext "timeline/backend/app/context"
	"timeline/backend/db/model/user"

	"github.com/rs/cors"
	"google.golang.org/api/idtoken"
)

func AuthMiddleware(googleClientID string, authorizeModel user.Authorize) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			token := extractToken(req)
			if token == "" {
				http.Error(w, "Empty token", http.StatusForbidden)
				return
			}

			payload, err := idtoken.Validate(context.Background(), token, googleClientID)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			someUser := user.NewSomeUser(payload.Claims["sub"].(string),
				payload.Claims["name"].(string),
				payload.Claims["email"].(string),
				payload.Claims["picture"].(string),
			)

			user := authorizeModel.CheckOrCreate(someUser)
			if user == nil {
				http.Error(w, "Blocked", http.StatusForbidden)
				return
			}

			ctx := context.WithValue(req.Context(), appContext.AppContextUserKey{}, user.ID)

			req = req.WithContext(ctx)
			next.ServeHTTP(w, req)
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

func Cors(allowedOrigin string, debug bool) *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:     []string{allowedOrigin},
		AllowCredentials:   true,
		AllowedMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:     []string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept", "Authorization"},
		OptionsPassthrough: true,
		Debug:              debug,
	})
}
