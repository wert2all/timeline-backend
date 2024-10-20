package context

import (
	"context"
)

type (
	userKey      struct{}
	userIsNewKey struct{}
)

func SetUserID(ctx context.Context, id int, isNew bool) context.Context {
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
