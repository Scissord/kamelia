package handler

import (
	"backend/internal/modules/auth/dto"
	authService "backend/internal/modules/auth/service"
	"backend/internal/utils"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var user dto.LoginUserDTO

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
	userProfile, err := h.service.Login(context.Background(), r, &user)
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
