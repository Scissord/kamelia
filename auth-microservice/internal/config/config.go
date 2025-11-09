package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	DBUrl string
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("⚠️ No .env file found, using system environment variables")
	}

	return &Config{
		Port:  os.Getenv("PORT"),
		DBUrl: os.Getenv("DATABASE_URL"),
	}
}
