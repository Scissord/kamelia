package main

import (
	"context"
	"fmt"
	"net/http"
	"task-app/internal/router"

	"github.com/jackc/pgx/v5"
)

const PORT = ":8080"

func main() {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:322434@localhost:5432/kamelia")
	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}
	defer conn.Close(context.Background())

	err = conn.Ping(context.Background())
	if err != nil {
		fmt.Println("Database ping error:", err)
		return
	}

	fmt.Println("Connected to PostgreSQL")

	r := router.InitRouter(conn)

	fmt.Println("Server started on port 8080")
	err = http.ListenAndServe(PORT, r)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
