package main

import (
	"log"
	"net/http"

	"backend/internal/config"
	"backend/internal/migrations"
	"backend/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	db := cfg.DB
	if err := migrations.AutoMigrateAll(db); err != nil {
		log.Fatal("Migration failed:", err)
	}

	srv := server.New(cfg)

	log.Println("Starting server on", cfg.HTTP.Addr)
	if err := http.ListenAndServe(cfg.HTTP.Addr, srv.Router); err != nil {
		log.Fatal(err)
	}
}
