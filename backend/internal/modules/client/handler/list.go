package handler

import (
	"backend/internal/modules/client/dto"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	var q dto.GetClientsQuery

	if v := r.URL.Query().Get("limit"); v != "" {
		q.Limit, _ = strconv.Atoi(v)
	}

	// Парсим page
	if v := r.URL.Query().Get("page"); v != "" {
		q.Page, _ = strconv.Atoi(v)
	}

	// Парсим sort
	if v := r.URL.Query().Get("sort"); v != "" {
		q.Sort = &v
	}

	if err := validate.Struct(q); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	clients, err := h.service.List(context.Background(), q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(clients); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
