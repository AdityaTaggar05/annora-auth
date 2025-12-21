package config

import (
	"os"
	"strconv"
)

type Config struct {
	PORT string
	DB_URL     string
	JWT_SECRET string
	JWT_EXP    int
}

func Load() Config {
	jwt_exp, _ := strconv.Atoi(os.Getenv("JWT_EXP"))

	return Config{
		PORT: os.Getenv("PORT"),
		DB_URL: os.Getenv("DB_URL"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
		JWT_EXP: jwt_exp,
	}
}