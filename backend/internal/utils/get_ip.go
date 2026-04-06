package utils

import (
	"net"
	"net/http"
	"strings"
)

func GetIP(r *http.Request) string {
	// 1. X-Forwarded-For (может быть список)
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0])
	}

	// 2. X-Real-IP (nginx)
	if xrip := r.Header.Get("X-Real-IP"); xrip != "" {
		return xrip
	}

	// 3. Fallback
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		return host
	}

	return r.RemoteAddr
}
