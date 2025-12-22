package models

import "time"

type RefreshToken struct {
	UserID    string
	Token     string
	Revoked   bool
	ExpiresAt time.Time
}
