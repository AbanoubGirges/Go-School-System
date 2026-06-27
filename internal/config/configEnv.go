package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DBUrl     string
	JwtSecret string
}

func Load() Config {
	err:=godotenv.Load()
	if err!=nil{
		panic("Failed to read env")
	}
	return Config{
		Port: os.Getenv("PORT"),
		DBUrl: os.Getenv("DATABASE_URL"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
}