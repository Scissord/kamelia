package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	DBUrl string
}

func Load() *Config {
	envPath := filepath.Join("..", "..", ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	return &Config{
		Port:  os.Getenv("PORT"),
		DBUrl: os.Getenv("DATABASE_URL"),
	}
}
