package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

// RequestID – обертка над chi middleware для request ID
var RequestID = middleware.RequestID

// Logger – простое логирование запроса
var Logger = middleware.Logger

// Recoverer – стандартный recover для chi
var Recoverer = middleware.Recoverer

// Timeout – таймаут для всех запросов (секунды)
func Timeout(seconds int) func(next http.Handler) http.Handler {
	return middleware.Timeout(time.Duration(seconds) * time.Second)
}

// ExampleAuth – пример кастомного middleware
func ExampleAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer secret" { // простая проверка
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

// SimpleLogger – пример собственного логгера
func SimpleLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}
