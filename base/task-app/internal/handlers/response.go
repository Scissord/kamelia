package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)

	if _, err := w.Write(bytes); err != nil {
		fmt.Println("write error:", err)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{
		"error": message,
	})
}
