package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"backend/internal/modules/client/dto"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var c dto.CreateClientDTO

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client, err := h.service.Create(context.Background(), &c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client)
}
