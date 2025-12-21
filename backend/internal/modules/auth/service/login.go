package service

import (
	"backend/internal/modules/auth/dto"
	userModel "backend/internal/modules/user/model"
	utils "backend/internal/utils"
	"context"
	"errors"
)

var ErrInvalidCredentials = errors.New("Invalid credentials")

func (s *service) Login(
	ctx context.Context,
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

	return &userModel.User{
		ID:        user.ID,
		Login:     user.Login,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// Если ты создаёшь сессию, убедись, что она уникальна для устройства.
// Можно хранить UserAgent + IP + expires_at.
// 4. Create session in db - auth.session
// 6. Generate pair of tokens
// 7. Create tokens in db - auth.token
// 8. Set refresh_token to cookies
// 9. Take profile - auth.profile.user_id
// 10. Don't forget to delete password_hash
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
