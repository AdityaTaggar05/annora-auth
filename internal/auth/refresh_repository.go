package auth

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type RefreshTokenRepository struct {
	DB *pgxpool.Pool
}
