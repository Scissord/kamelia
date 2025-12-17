package middleware

import (
	"net/http"
	"time"

	chimw "github.com/go-chi/chi/v5/middleware"
)

var (
	RequestID = chimw.RequestID
	Logger    = chimw.Logger
	Recoverer = chimw.Recoverer
)

func Timeout(seconds int) func(http.Handler) http.Handler {
	return chimw.Timeout(time.Duration(seconds) * time.Second)
}
