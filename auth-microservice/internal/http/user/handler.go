package user

import (
	"encoding/json"
	"net/http"

	types "auth-microservice/internal/schema/user"
	service "auth-microservice/internal/service/user"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	var input types.RegistrationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	user, err := h.service.Registration(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := types.RegistrationResponse{
		ID:    user.ID,
		Login: user.Login,
	}

	//TODO: if email, phone, birthday, locale, timezone is sended need to create auth.profile

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
