package context

import (
	"context"
)

type userKey struct{}

func SetUserID(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, userKey{}, id)
}

func GetUserID(ctx context.Context) int {
	val, _ := ctx.Value(userKey{}).(int)
	return val
}
