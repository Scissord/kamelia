package auth

import (
	"encoding/json"
	"log"
	"net/http"

	authTypes "auth-microservice/internal/schema/auth"

	profileRepo "auth-microservice/internal/repository/profile"

	profileService "auth-microservice/internal/service/profile"
	userService "auth-microservice/internal/service/user"
)

type Handler struct {
	userService    *userService.Service
	profileService *profileService.Service
}

func NewHandler(userSrv *userService.Service, profileSrv *profileService.Service) *Handler {
	return &Handler{
		userService:    userSrv,
		profileService: profileSrv,
	}
}

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	var input authTypes.RegistrationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	user, err := h.userService.Registration(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userJSON, _ := json.Marshal(user)
	log.Printf("User created: %s", userJSON)

	var profile *profileRepo.Profile
	if input.FirstName != nil || input.MiddleName != nil || input.LastName != nil ||
		input.Email != nil || input.Phone != nil || input.Birthday != nil ||
		input.Gender != nil || input.Locale != nil || input.Timezone != nil {

		profile, err = h.profileService.Create(user, &input)
		if err != nil {
			http.Error(w, "Failed to create profile: "+err.Error(), http.StatusInternalServerError)
			return
		}

		profileJSON, _ := json.Marshal(profile)
		log.Printf("Profile created: %s", profileJSON)
	}

	resp := authTypes.RegistrationResponse{
		ID:        user.ID,
		Login:     user.Login,
		CreatedAt: &user.CreatedAt,
	}

	if profile != nil {
		resp.FirstName = profile.FirstName
		resp.MiddleName = profile.MiddleName
		resp.LastName = profile.LastName
		resp.Email = profile.Email
		resp.Phone = profile.Phone
		resp.Birthday = profile.Birthday
		resp.Gender = &profile.Gender
		resp.Locale = &profile.Locale
		resp.Timezone = &profile.Timezone
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
