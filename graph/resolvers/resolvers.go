package resolvers

import (
	"context"
)

type Arguments[T any] interface {
	GetArguments() T
}

type ValidArguments[T any] interface {
	GetArguments() T
}

type Validator[K any, N any] interface {
	Validate(context.Context, Arguments[K]) (ValidArguments[N], error)
}
type Resolver[T any, K any] interface {
	Resolve(context.Context, ValidArguments[K]) (T, error)
}

func Resolve[T any, K any, N any](context context.Context, inputArguments Arguments[K], validator Validator[K, N], resolver Resolver[T, N]) (T, error) {
	var result T
	valid, err := validator.Validate(context, inputArguments)
	if err != nil {
		return result, err
	}
	return resolver.Resolve(context, valid)
}
