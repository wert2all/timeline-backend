package context

import (
	"context"
)

type (
	userKey      struct{}
	userIsNewKey struct{}
	tokenKey     struct{}
)

func SetUserID(ctx context.Context, id int, isNew bool, token string) context.Context {
	ctxWithIsNew := context.WithValue(ctx, userIsNewKey{}, isNew)
	ctxWithUser := context.WithValue(ctxWithIsNew, userKey{}, id)
	ctxWithToken := context.WithValue(ctxWithUser, tokenKey{}, token)

	return ctxWithToken
}

func GetUserID(ctx context.Context) int {
	val, _ := ctx.Value(userKey{}).(int)
	return val
}

func GetIsNew(ctx context.Context) bool {
	val, _ := ctx.Value(userIsNewKey{}).(bool)
	return val
}

func GetToken(ctx context.Context) string {
	val, _ := ctx.Value(tokenKey{}).(string)
	return val
}
