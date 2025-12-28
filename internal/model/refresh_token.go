package model

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

type RefreshToken struct {
	UserID    string
	Token     string
	Revoked   bool
	ExpiresAt time.Time
}

func GenerateRefreshToken(userID string, ttl time.Duration) (RefreshToken, error) {
	b := make([]byte, 32)

	_, err := rand.Read(b)
	if err != nil {
		return RefreshToken{}, nil
	}

	token := base64.URLEncoding.EncodeToString(b)

	return RefreshToken{
		UserID:    userID,
		Token:     token,
		Revoked:   false,
		ExpiresAt: time.Now().Add(ttl),
	},  nil
}
