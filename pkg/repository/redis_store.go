package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) *RedisStore {
	return &RedisStore{client: client}
}

func (r *RedisStore) SetSession(ctx context.Context, SID string, id int, lifetime time.Duration) error {
	if err := r.client.Set(ctx, SID, id, lifetime).Err(); err != nil {
		return err
	}
	return nil
}

func (r *RedisStore) GetSession(ctx context.Context, SID string) (int, error) {
	val, err := r.client.Get(ctx, SID).Result()
	if err != nil {
		return 0, err
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	return intVal, nil
}

func (r *RedisStore) DeleteSession(ctx context.Context, SID string) error {
	err := r.client.Del(ctx, SID).Err()
	if err != nil {
		return err
	}
	return nil
}
