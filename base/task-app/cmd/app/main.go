package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5"

	"task-app/internal/config"
	"task-app/internal/router"
)

func main() {
	cfg := config.Load()

	conn, err := pgx.Connect(context.Background(), cfg.DBURL)
	if err != nil {
		fmt.Println("database connection error:", err)
		return
	}
	fmt.Println("Connected to PostgreSQL")

	r := router.InitRouter(conn, cfg)

	server := &http.Server{
		Addr:    cfg.Port,
		Handler: r,
	}

	serverErrors := make(chan error, 1)

	go func() {
		fmt.Println("Server started on port", cfg.Port)

		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		fmt.Println("server error:", err)

	case sig := <-quit:
		fmt.Println("shutdown signal received:", sig)

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			fmt.Println("server shutdown error:", err)
		}

		if err := conn.Close(context.Background()); err != nil {
			fmt.Println("database close error:", err)
		}

		fmt.Println("server stopped gracefully")
	}
}
