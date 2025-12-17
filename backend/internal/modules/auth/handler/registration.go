package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"backend/internal/modules/auth/dto"
)

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	var u dto.RegistrationUserDTO

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.service.Registration(context.Background(), &u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
