package redisext

import (
	"context"
	"fmt"

	"github.com/ThisJohan/snapp-assignment/pkg/di"
	"github.com/redis/go-redis/v9"
)

const Service = "redis"

type Config struct {
	RedisPort string `env:"REDIS_PORT" envDefault:"6379"`
	RedisHost string `env:"REDIS_HOST" envDefault:"localhost"`
}

func Connect(config Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
		Password: "",
		DB:       0,
	})
	return rdb
}

func InjectRedis(ctx context.Context) *redis.Client {
	return di.Inject[*redis.Client](ctx, Service)
}
