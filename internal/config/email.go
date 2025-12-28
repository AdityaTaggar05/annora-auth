package config

import "time"

type EmailConfig struct {
	From     string
	SMTPHost string
	SMTPPort int
	Username string
	Password string
	TokenTTL time.Duration
}