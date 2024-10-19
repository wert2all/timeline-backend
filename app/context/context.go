package context

import (
	"context"
)

type userKey struct{}
type userIsNewKey struct{}

func GetUserID(ctx context.Context) int {
	val, _ := ctx.Value(userKey{}).(int)
	return val
}

func GetIsNew(ctx context.Context) bool {
	val, _ := ctx.Value(userIsNewKey{}).(bool)
	return val
}
