package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id, secret string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":id,
		"exp":time.Now().Add(ttl).Unix(),
	})

	return token.SignedString([]byte(secret))
}