package rediscli

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/zjrb/OpeningTrainer/internal/core/domain"
)

type RedisCache struct {
	cli *redis.Client
}

func NewRedisRepo(rediscli *redis.Client) *RedisCache {
	return &RedisCache{cli: rediscli}
}

func (r *RedisCache) SetOpening(key string, opening *domain.GameSesion) error {
	if _, err := r.cli.Pipelined(context.Background(), func(rdb redis.Pipeliner) error {
		rdb.HSet(context.Background(), key, "openining", opening.Opening)
		rdb.HSet(context.Background(), key, "white", opening.White)
		rdb.HSet(context.Background(), key, "moveNum", opening.MoveNum)
		rdb.HSet(context.Background(), key, "lastMove", opening.LastMove)
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (r *RedisCache) GetOpening(key string) (*domain.GameSesion, error) {
	var gameSesh domain.GameSesion
	if err := r.cli.HGetAll(context.Background(), key).Scan(&gameSesh); err != nil {
		return nil, err
	}
	return &gameSesh, nil
}
