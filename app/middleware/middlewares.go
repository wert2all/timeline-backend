package middlewares

import (
	"context"
	"net/http"
	"strings"

	"timeline/backend/db/model/user"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"google.golang.org/api/idtoken"
)

type (
	Middlewares struct {
		List []func(http.Handler) http.Handler
	}
	userKey      struct{}
	userIsNewKey struct{}
)

func NewMiddlewares(userModel user.Authorize, googleClientID string) Middlewares {
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
			authMiddleware(userModel, googleClientID),
		},
	}
}

func authMiddleware(userModel user.Authorize, googleClientID string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// fmt.Printf("User ID: %d\n", context.GetUserID())

			token := extractToken(req)
			if token == "" {
				http.Error(w, "Empty token", http.StatusForbidden)
				return
			}

			payload, err := idtoken.Validate(req.Context(), token, googleClientID)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			someUser := user.NewSomeUser(payload.Claims["sub"].(string),
				payload.Claims["name"].(string),
				payload.Claims["email"].(string),
				payload.Claims["picture"].(string),
			)

			userCheck, error := userModel.CheckOrCreate(someUser)
			if error != nil {
				http.Error(w, "Blocked", http.StatusForbidden)
				return
			}

			req = req.WithContext(setUserID(req.Context(), userCheck.ID, userCheck.IsNew))
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

func setUserID(ctx context.Context, id int, isNew bool) context.Context {
	newCtx := context.WithValue(ctx, userIsNewKey{}, isNew)
	return context.WithValue(newCtx, userKey{}, id)
}

func GetUserID(ctx context.Context) int {
	val, _ := ctx.Value(userKey{}).(int)
	return val
}

func GetIsNew(ctx context.Context) bool {
	val, _ := ctx.Value(userIsNewKey{}).(bool)
	return val
}