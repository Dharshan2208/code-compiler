package limiter

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisManager struct {
	client *redis.Client

	capacity   float64
	refillRate float64
}

func NewRedisManager(client *redis.Client, capacity float64, refillRate float64) *RedisManager {
	return &RedisManager{
		client:     client,
		capacity:   capacity,
		refillRate: refillRate,
	}
}

func (m *RedisManager) Allow(key string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	redisKey := "ratelimit:" + key

	result, err := allowScript.Run(
		ctx,
		m.client,
		[]string{redisKey},
		m.capacity,
		m.refillRate,
		time.Now().UnixMilli(),
		1800,
	).Int()
	if err != nil {
		return false
	}

	return result == 1
}
