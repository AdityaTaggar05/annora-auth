package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	PORT string
	DB_URL     string
	JWT_SECRET string
	JWT_EXP    time.Duration
	REFRESH_EXP time.Duration
}

func Load() Config {
	jwt_exp, _ := strconv.Atoi(os.Getenv("JWT_EXP"))
	refresh_exp, _ := strconv.Atoi(os.Getenv("REFRESH_EXP"))

	return Config{
		PORT: os.Getenv("PORT"),
		DB_URL: os.Getenv("DATABASE_URL"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
		JWT_EXP: time.Duration(jwt_exp) * time.Minute,
		REFRESH_EXP: time.Duration(refresh_exp) * (time.Hour * 24),
	}
}