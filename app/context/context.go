package context

import (
	"context"
)

type (
	tokenKey struct{}
)

func SetUserID(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenKey{}, token)
}

func GetToken(ctx context.Context) string {
	val, _ := ctx.Value(tokenKey{}).(string)
	return val
}
