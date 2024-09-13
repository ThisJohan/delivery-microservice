package di

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

type diKey string

func GrpcProvide(builders ...Builder) grpc.ServerOption {
	return grpc.UnaryInterceptor(func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp any, err error) {
		for _, builder := range builders {
			ctx = builder(ctx)
		}
		return handler(ctx, req)
	})
}

func FromContext(ctx context.Context, key string, val any) context.Context {
	return context.WithValue(ctx, diKey(key), val)
}

func Inject[T any](ctx context.Context, key string) T {
	value, ok := ctx.Value(diKey(key)).(T)
	if !ok {
		log.Fatalf("DI: Failed to get value from context with key: %s", key)
	}
	return value
}

type Builder func(context.Context) context.Context

func DIBuilder[T any](key string, provide func() T) Builder {
	value := provide()
	return func(ctx context.Context) context.Context {
		return FromContext(ctx, key, value)
	}
}
