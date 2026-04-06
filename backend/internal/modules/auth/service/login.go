package service

import (
	"backend/internal/modules/auth/dto"
	sessionModel "backend/internal/modules/session/model"
	userModel "backend/internal/modules/user/model"
	utils "backend/internal/utils"
	"context"
	"errors"
	"net/http"
	"time"
)

var ErrInvalidCredentials = errors.New("Invalid credentials")
var ErrInvalidSession = errors.New("Error while creating session")

func (s *service) Login(
	ctx context.Context,
	r *http.Request,
	dto *dto.LoginUserDTO,
) (*userModel.User, error) {
	// 1. Check if login in base
	user, err := s.repo.GetByLogin(ctx, *dto.Login)
	if err != nil {
		// пользователь не найден или ошибка БД
		return nil, ErrInvalidCredentials
	}

	// 2. Check if password compare
	if !utils.CheckPasswordHash(dto.Password, user.PasswordHash) {
		return nil, ErrInvalidCredentials
	}

	// Если ты создаёшь сессию, убедись, что она уникальна для устройства.
	// Можно хранить UserAgent + IP + expires_at.
	// 3. Create session in db - auth.session
	now := time.Now()

	session_input := &sessionModel.Session{
		UserID:     user.ID,
		IPAddress:  utils.GetIP(r),
		UserAgent:  r.UserAgent(),
		LoginAt:    now,
		LogoutAt:   nil,
		LastSeenAt: &now,
		IsActive:   true,
	}
	session_output, err := s.repo.CreateSession(ctx, session_input)
	if err != nil {
		return nil, ErrInvalidSession
	}

	return &userModel.User{
		ID:        user.ID,
		Login:     user.Login,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// 4. Generate pair of tokens
// 5. Create tokens in db - auth.token
// 6. Set refresh_token to cookies
// 7. Take profile - auth.profile.user_id
// 8. Don't forget to delete password_hash
// return {
// 	user: {
// 		...user,
// 		...profile,
// 		...session,
// 		access_token,
// 	}
// }

// Audit / логирование
// Каждое логирование login/logout.
// IP, device, timestamp → помогает обнаруживать аномалии.
// }
