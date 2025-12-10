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
	// Загружаем .env (если нет — не ошибка)
	_ = godotenv.Load()

	// Получаем ENV
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	httpAddr := os.Getenv("HTTP_ADDR")

	// Формируем DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbHost, dbUser, dbPass, dbName, dbPort,
	)

	// Подключение через GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		DB: db,
		HTTP: HTTPConfig{
			Addr: httpAddr,
		},
	}

	return cfg, nil
}
