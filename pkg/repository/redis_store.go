package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) *RedisStore {
	return &RedisStore{client: client}
}

func (r *RedisStore) SetSession(SID string, id int) error {
	if err := r.client.Set(context.Background(), SID, id, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (r *RedisStore) GetSession(SID string) (int, error) {
	val, err := r.client.Get(context.Background(), SID).Int()
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (r *RedisStore) DeleteSession(SID string) error {
	err := r.client.Del(context.Background(), SID).Err()
	if err != nil {
		return err
	}
	return nil
}
