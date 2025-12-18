package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/modules/client/dto"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)

	var c dto.UpdateClientDTO
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client, err := h.service.Update(context.Background(), uint(id), &c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client)
}
