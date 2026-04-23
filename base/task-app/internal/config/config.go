package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	DBURL string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Ошибка загрузки .env:", err)
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is required")
	}

	return &Config{
		Port:  ":" + port,
		DBURL: dbURL,
	}
}
