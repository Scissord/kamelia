package utils

import (
	"encoding/json"
	"net/http"

	model "backend/internal/modules/auth/model"
)

func WriteJSONError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(model.ErrorResponse{Error: msg})
}
