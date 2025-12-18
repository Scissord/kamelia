package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config хранит настройки приложения
type Config struct {
	DB   *gorm.DB
	HTTP HTTPConfig
}

// HTTPConfig хранит настройки HTTP-сервера
type HTTPConfig struct {
	Addr string
}

// Load создаёт конфиг с подключением к БД через GORM
func Load() (*Config, error) {
	// Находим .env файл относительно текущего файла или рабочей директории
	// Пробуем несколько путей
	envPaths := []string{
		".env",
		"../.env",
		"../../.env",
		"./backend/.env",
	}

	for _, path := range envPaths {
		if err := godotenv.Load(path); err == nil {
			break
		}
	}

	// Получаем ENV
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	httpAddr := os.Getenv("HTTP_ADDR")

	if httpAddr == "" {
		httpAddr = ":8080"
	}

	// Формируем DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbHost, dbUser, dbPass, dbName, dbPort,
	)

	// Подключение через GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	cfg := &Config{
		DB: db,
		HTTP: HTTPConfig{
			Addr: httpAddr,
		},
	}

	return cfg, nil
}
