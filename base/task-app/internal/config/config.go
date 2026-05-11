package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                  string
	DBURL                 string
	JWTSecret             string
	AccessTokenTTLMinutes int
	RefreshTokenTTLDays   int
	SecureCookie          bool
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

	jwtSECRET := os.Getenv("JWT_SECRET")
	if jwtSECRET == "" {
		log.Fatal("JWT_SECRET is required")
	}

	accessTokenTTLMinutes := 15

	if value := os.Getenv("ACCESS_TOKEN_TTL_MINUTES"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("invalid ACCESS_TOKEN_TTL_MINUTES")
		}

		accessTokenTTLMinutes = parsed
	}

	refreshTokenTTLDays := 30

	if value := os.Getenv("REFRESH_TOKEN_TTL_DAYS"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("invalid REFRESH_TOKEN_TTL_DAYS")
		}

		refreshTokenTTLDays = parsed
	}

	secureCookie := false

	if value := os.Getenv("SECURE_COOKIE"); value != "" {
		parsed, err := strconv.ParseBool(value)
		if err != nil {
			log.Fatal("invalid SECURE_COOKIE")
		}

		secureCookie = parsed
	}

	return &Config{
		Port:                  ":" + port,
		DBURL:                 dbURL,
		JWTSecret:             jwtSECRET,
		AccessTokenTTLMinutes: accessTokenTTLMinutes,
		RefreshTokenTTLDays:   refreshTokenTTLDays,
		SecureCookie:          secureCookie,
	}
}
