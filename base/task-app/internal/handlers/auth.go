package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"

	"task-app/internal/auth"
	"task-app/internal/config"
	m2 "task-app/internal/middleware"
	"task-app/internal/models"
	"task-app/internal/repository"
)

type AuthHandler struct {
	UserRepo  *repository.UserRepository
	TokenRepo *repository.TokenRepository
	Config    *config.Config
}

func NewAuthHandler(
	userRepo *repository.UserRepository,
	tokenRepo *repository.TokenRepository,
	cfg *config.Config,
) *AuthHandler {
	return &AuthHandler{
		UserRepo:  userRepo,
		TokenRepo: tokenRepo,
		Config:    cfg,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var input models.RegisterInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	input.Email = strings.TrimSpace(strings.ToLower(input.Email))

	if input.Email == "" {
		writeError(w, http.StatusBadRequest, "email is required")
		return
	}

	if len(input.Password) < 8 {
		writeError(w, http.StatusBadRequest, "password must be at least 8 characters")
		return
	}

	_, _, err = h.UserRepo.GetUserByEmail(input.Email)
	if err == nil {
		writeError(w, http.StatusConflict, "email already exists")
		return
	}

	if err != pgx.ErrNoRows {
		writeError(w, http.StatusInternalServerError, "failed to check user")
		return
	}

	passwordHashBytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to hash password")
		return
	}

	user, err := h.UserRepo.CreateUser(input.Email, string(passwordHashBytes))
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create user")
		return
	}

	writeJSON(w, http.StatusCreated, user)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var input models.LoginInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	input.Email = strings.TrimSpace(strings.ToLower(input.Email))

	if input.Email == "" {
		writeError(w, http.StatusBadRequest, "email is required")
		return
	}

	if input.Password == "" {
		writeError(w, http.StatusBadRequest, "password is required")
		return
	}

	user, passwordHash, err := h.UserRepo.GetUserByEmail(input.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			writeError(w, http.StatusUnauthorized, "invalid email or password")
			return
		}

		writeError(w, http.StatusInternalServerError, "failed to get user")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(input.Password))
	if err != nil {
		writeError(w, http.StatusUnauthorized, "invalid email or password")
		return
	}

	accessToken, err := auth.GenerateToken(
		user.ID,
		user.Email,
		h.Config.JWTSecret,
		time.Duration(h.Config.AccessTokenTTLMinutes)*time.Minute,
	)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create access token")
		return
	}

	refreshToken, err := auth.GenerateToken(
		user.ID,
		user.Email,
		h.Config.JWTSecret,
		time.Duration(h.Config.RefreshTokenTTLDays)*24*time.Hour,
	)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create refresh token")
		return
	}

	refreshTokenHash := auth.HashToken(refreshToken)

	refreshTokenExpiresAt := time.Now().Add(
		time.Duration(h.Config.RefreshTokenTTLDays) * 24 * time.Hour,
	)

	err = h.TokenRepo.CreateRefreshToken(
		user.ID,
		refreshTokenHash,
		refreshTokenExpiresAt,
	)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to save refresh token")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   h.Config.SecureCookie,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   h.Config.RefreshTokenTTLDays * 24 * 60 * 60,
	})

	writeJSON(w, http.StatusOK, models.AuthResponse{
		User:        user,
		AccessToken: accessToken,
	})
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		writeError(w, http.StatusUnauthorized, "refresh token is missing")
		return
	}

	refreshToken := cookie.Value
	refreshTokenHash := auth.HashToken(refreshToken)

	storedToken, err := h.TokenRepo.GetByHash(refreshTokenHash)
	if err != nil {
		if err == pgx.ErrNoRows {
			writeError(w, http.StatusUnauthorized, "invalid refresh token")
			return
		}

		writeError(w, http.StatusInternalServerError, "failed to get refresh token")
		return
	}

	if storedToken.RevokedAt != nil {
		writeError(w, http.StatusUnauthorized, "refresh token revoked")
		return
	}

	if time.Now().After(storedToken.ExpiresAt) {
		writeError(w, http.StatusUnauthorized, "refresh token expired")
		return
	}

	err = h.TokenRepo.RevokeToken(storedToken.ID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to revoke refresh token")
		return
	}

	user, err := h.UserRepo.GetUserByID(storedToken.UserID)
	if err != nil {
		if err == pgx.ErrNoRows {
			writeError(w, http.StatusUnauthorized, "user not found")
			return
		}

		writeError(w, http.StatusInternalServerError, "failed to get user")
		return
	}

	accessToken, err := auth.GenerateToken(
		user.ID,
		user.Email,
		h.Config.JWTSecret,
		time.Duration(h.Config.AccessTokenTTLMinutes)*time.Minute,
	)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create access token")
		return
	}

	newRefreshToken, err := auth.GenerateToken(
		user.ID,
		user.Email,
		h.Config.JWTSecret,
		time.Duration(h.Config.RefreshTokenTTLDays)*24*time.Hour,
	)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create refresh token")
		return
	}

	newRefreshTokenHash := auth.HashToken(newRefreshToken)
	newRefreshTokenExpiresAt := time.Now().Add(
		time.Duration(h.Config.RefreshTokenTTLDays) * 24 * time.Hour,
	)

	err = h.TokenRepo.CreateRefreshToken(
		user.ID,
		newRefreshTokenHash,
		newRefreshTokenExpiresAt,
	)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to save refresh token")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    newRefreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   h.Config.SecureCookie,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   h.Config.RefreshTokenTTLDays * 24 * 60 * 60,
	})

	writeJSON(w, http.StatusOK, models.AuthResponse{
		User:        user,
		AccessToken: accessToken,
	})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		writeError(w, http.StatusUnauthorized, "refresh token is missing")
		return
	}

	refreshTokenHash := auth.HashToken(cookie.Value)

	storedToken, err := h.TokenRepo.GetByHash(refreshTokenHash)
	if err != nil {
		if err == pgx.ErrNoRows {
			writeError(w, http.StatusUnauthorized, "invalid refresh token")
			return
		}

		writeError(w, http.StatusInternalServerError, "failed to get refresh token")
		return
	}

	if storedToken.RevokedAt == nil {
		err = h.TokenRepo.RevokeToken(storedToken.ID)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "failed to logout")
			return
		}
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   h.Config.SecureCookie,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	})

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "logged out",
	})
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(m2.UserContextKey).(*auth.TokenClaims)
	if !ok {
		writeError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	user, err := h.UserRepo.GetUserByID(claims.UserID)
	if err != nil {
		if err == pgx.ErrNoRows {
			writeError(w, http.StatusUnauthorized, "user not found")
			return
		}

		writeError(w, http.StatusInternalServerError, "failed to get user")
		return
	}

	writeJSON(w, http.StatusOK, user)
}
