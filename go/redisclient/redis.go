package redisclient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	client *redis.Client
}

func Open() *RedisService {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	return &RedisService{client}
}

func (rs *RedisService) Ping(ctx context.Context) (string, error) {
	return rs.client.Ping(ctx).Result()
}
