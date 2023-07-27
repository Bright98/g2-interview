package repository

import (
	"context"
	"g2/auth/sso/domain"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
}

func NewRepository() domain.RepositoryInterface {
	repo := &Repository{}
	return repo
}

// redis
func RedisConnection(address, password string, db int) error {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	RedisClient = client
	RedisCtx = context.Background()

	return nil
}
