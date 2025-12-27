package utils

import (
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id string, privateKey *rsa.PrivateKey, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.RegisteredClaims{
		Subject: id,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
		IssuedAt: jwt.NewNumericDate(time.Now()),
	})

	return token.SignedString(privateKey)
}