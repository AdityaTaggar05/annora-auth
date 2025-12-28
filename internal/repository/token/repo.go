package tokenrepo

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type TokenRepository struct {
	DB *pgxpool.Pool
	RDB *redis.Client
}

func NewRepository(db *pgxpool.Pool, rdb *redis.Client) *TokenRepository {
	return &TokenRepository{
		DB: db,
		RDB: rdb,
	}
}
