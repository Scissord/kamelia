package repository

import (
	"backend/internal/modules/profile/model"
	"backend/internal/utils"
	"context"
	"os"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	createProfile(ctx context.Context, tx *gorm.DB, profile *model.Profile) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) createProfile(
	ctx context.Context,
	tx *gorm.DB,
	profile *model.Profile,
) error {

	key := os.Getenv("PGP_SECRET")

	data := map[string]interface{}{
		"user_id": profile.UserID,

		"first_name_encrypted":  utils.Encrypt(profile.FirstName, key),
		"last_name_encrypted":   utils.Encrypt(profile.LastName, key),
		"middle_name_encrypted": utils.Encrypt(profile.MiddleName, key),
		"email_encrypted":       utils.Encrypt(profile.Email, key),
		"phone_encrypted":       utils.Encrypt(profile.Phone, key),
		"birthday_encrypted":    utils.Encrypt(profile.Birthday, key),

		"gender":   profile.Gender,
		"locale":   profile.Locale,
		"timezone": profile.Timezone,
	}

	return tx.WithContext(ctx).
		Model(&model.Profile{}).
		Create(data).Error
}
