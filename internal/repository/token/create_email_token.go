package tokenrepo

import (
	"context"
	"time"
)

func (r *TokenRepository) CreateEmailToken(ctx context.Context, key, userID string, ttl time.Duration) error {
	return r.RDB.Set(ctx, key, userID, ttl).Err()
}