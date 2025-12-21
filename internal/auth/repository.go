package auth

import (
	"context"

	"github.com/AdityaTaggar05/connectify-auth/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	DB *pgxpool.Pool
}

func (r *Repository) CreateUser(ctx context.Context, email, hash string) error {
	_, err := r.DB.Exec(ctx,
	`INSERT INTO users (email, password_hash) VALUES ($1, $2)`,
	email, hash)

	return err
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User

	err := r.DB.QueryRow(ctx,
	`SELECT id, email, password_hash, created_at FROM users WHERE email=$1`,
	email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)

	return user, err
}