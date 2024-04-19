package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"
	"timeline/backend/db/model/user"
	"timeline/backend/http/middleware/auth"

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

			user := authorizeModel.CheckOrCreate(auth.From(*payload))
			if user == nil {
				http.Error(w, "Blocked", http.StatusForbidden)
				return
			}

			log.Println("user was created: ", user)

			//
			// // Allow unauthenticated users in
			// if err != nil || c == nil {
			// 	next.ServeHTTP(w, r)
			// 	return
			// }

			// userId, err := validateAndGetUserID(c)
			// if err != nil {
			// 	http.Error(w, "Invalid cookie", http.StatusForbidden)
			// 	return
			// }

			// // get the user from the database
			// user := getUserByID(db, userId)

			// // put it in context
			// ctx := context.WithValue(r.Context(), userCtxKey, user)

			// // and call the next with our new context
			// r = r.WithContext(ctx)
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
