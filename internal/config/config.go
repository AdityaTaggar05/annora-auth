package config

import "time"

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	JWT      JWTConfig
	Redis    RedisConfig
	Email    EmailConfig
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port:         getEnv("PORT", "8080"), 
			ReadTimeout:  getDuration("SERVER_READ_TIMEOUT", 5*time.Second),
			WriteTimeout: getDuration("SERVER_WRITE_TIMEOUT", 10*time.Second),
		},

		Postgres: PostgresConfig{
			URL:          mustGetEnv("DATABASE_URL"),
			MaxOpenConns: getInt("DB_MAX_OPEN_CONNS", 10),
		},

		JWT: JWTConfig{
			PrivateKeyPath: mustGetEnv("JWT_PRIVATE_KEY_PATH"),
			PublicKeyPath:  mustGetEnv("JWT_PUBLIC_KEY_PATH"),
			Issuer:         getEnv("JWT_ISSUER", "annora-auth"),
			AccessTTL:      getDuration("JWT_ACCESS_TTL", 15*time.Minute),
			RefreshTTL:     getDuration("JWT_REFRESH_TTL", 7*24*time.Hour),
		},

		Email: EmailConfig{
			From:     getEnv("EMAIL_FROM", ""),
			SMTPHost: getEnv("SMTP_HOST", ""),
			SMTPPort: getInt("SMTP_PORT", 587),
			Username: getEnv("SMTP_USERNAME", ""),
			Password: getEnv("SMTP_PASSWORD", ""),
			TokenTTL: getDuration("EMAIL_TOKEN_TTL", 4*time.Hour),
		},
	}
}
