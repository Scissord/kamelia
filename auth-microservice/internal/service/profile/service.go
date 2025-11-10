package profile

import (
	repo "auth-microservice/internal/repository/profile"
	userRepo "auth-microservice/internal/repository/user"
	authTypes "auth-microservice/internal/schema/auth"
)

type Service struct {
	repo *repo.Repository
}

func NewService(r *repo.Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(user *userRepo.User, input *authTypes.RegistrationInput) (*repo.Profile, error) {
	// если все поля пустые — ничего не делаем
	if input.FirstName == nil && input.LastName == nil && input.Email == nil &&
		input.Phone == nil && input.Birthday == nil && input.Gender == nil &&
		input.Locale == nil && input.Timezone == nil {
		return nil, nil
	}

	profile := &repo.ProfileCreateInput{
		UserID:     user.ID,
		FirstName:  input.FirstName,
		MiddleName: input.MiddleName,
		LastName:   input.LastName,
		Email:      input.Email,
		Phone:      input.Phone,
		Birthday:   input.Birthday,
		Gender:     input.Gender,
		Locale:     input.Locale,
		Timezone:   input.Timezone,
	}

	return s.repo.Create(profile)
}
