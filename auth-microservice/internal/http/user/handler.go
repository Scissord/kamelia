package user

import (
	"encoding/json"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

	repo "auth-microservice/internal/repository/user"
	service "auth-microservice/internal/service/user"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

type RegistrationInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegistrationResponse struct {
	ID    int64  `json:"id"`
	Login string `json:"login"`
}

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	// 1. validate data for login, password, trim
	var input RegistrationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// 2. Trim and basic validate
	input.Login = strings.TrimSpace(input.Login)
	input.Password = strings.TrimSpace(input.Password)

	if input.Login == "" || input.Password == "" {
		http.Error(w, "Login and password are required", http.StatusBadRequest)
		return
	}
	if len(input.Password) < 6 {
		http.Error(w, "Password must be at least 6 characters", http.StatusBadRequest)
		return
	}

	// 3. try to find by login
	if err := h.service.FindByLogin(input.Login); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. hash password in js i used bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// 5. create user
	user := &repo.User{
		Login:        input.Login,
		PasswordHash: string(hashedPassword),
	}

	if err := h.service.Create(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := RegistrationResponse{
		ID:    user.ID,
		Login: user.Login,
	}

	//TODO: if email, phone, birthday, locale, timezone is sended need to create auth.profile

	// 6. send data to frontend
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
