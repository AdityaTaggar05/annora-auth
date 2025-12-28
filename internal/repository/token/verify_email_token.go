package tokenrepo

import (
	"context"
)

func (r *TokenRepository) VerifyEmailToken(ctx context.Context, key string) (string, error) {
	userID, err := r.RDB.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	r.RDB.Del(ctx, key)
	return userID, nil
}