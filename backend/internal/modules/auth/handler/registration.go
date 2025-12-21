package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"backend/internal/modules/auth/dto"
	"backend/internal/utils"

	authService "backend/internal/modules/auth/service"
)

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	var user dto.RegistrationUserDTO

	// Decoding body from request
	// If empty return EOF
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate income data by dto
	if err := validate.Struct(user); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Validation error")
		return
	}

	// Call service for logic with db
	userProfile, err := h.service.Registration(context.Background(), &user)
	if err != nil {
		log.Println("Registration service error:", err)
		switch {
		case errors.Is(err, authService.ErrEmailExists):
			utils.WriteJSONError(w, http.StatusConflict, "User already exists")
		case errors.Is(err, authService.ErrPhoneExists):
			utils.WriteJSONError(w, http.StatusConflict, "Phone already exists")
		default:
			utils.WriteJSONError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(userProfile); err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Internal Server Error")
		return
	}
}
