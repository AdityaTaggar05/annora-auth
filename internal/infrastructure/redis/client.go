package redisinfra

import (
	"github.com/AdityaTaggar05/annora-auth/internal/config"
	"github.com/redis/go-redis/v9"
)


func NewClient(cfg config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
		Password: cfg.Password,
		DB: cfg.DB,
	})
}
