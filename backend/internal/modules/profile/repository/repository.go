package repository

import (
	profileModel "backend/internal/modules/profile/model"
	"backend/internal/utils"
	"context"
	"os"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	createProfile(ctx context.Context, tx *gorm.DB, profile *profileModel.Profile) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) createProfile(
	ctx context.Context,
	tx *gorm.DB,
	profile *profileModel.Profile,
) error {
	key := os.Getenv("PGP_SECRET")

	// Считаем хеши
	profile.FirstNameHash = utils.Hash(profile.FirstName)
	profile.LastNameHash = utils.Hash(profile.LastName)
	profile.MiddleNameHash = utils.Hash(profile.MiddleName)
	profile.EmailHash = utils.Hash(profile.Email)
	profile.PhoneHash = utils.Hash(profile.Phone)
	profile.BirthdayHash = utils.Hash(profile.Birthday)

	// Сохраняем объект напрямую
	err := tx.WithContext(ctx).Create(profile).Error
	if err != nil {
		return err
	}

	// После Create, ID автоматически присваивается profile.ID

	// Затем обновляем поля с шифрованием
	return tx.WithContext(ctx).Model(profile).Updates(map[string]interface{}{
		"first_name_encrypted":  gorm.Expr("pgp_sym_encrypt(?, ?)", *profile.FirstName, key),
		"last_name_encrypted":   gorm.Expr("pgp_sym_encrypt(?, ?)", *profile.LastName, key),
		"middle_name_encrypted": gorm.Expr("pgp_sym_encrypt(?, ?)", *profile.MiddleName, key),
		"email_encrypted":       gorm.Expr("pgp_sym_encrypt(?, ?)", *profile.Email, key),
		"phone_encrypted":       gorm.Expr("pgp_sym_encrypt(?, ?)", *profile.Phone, key),
		"birthday_encrypted":    gorm.Expr("pgp_sym_encrypt(?, ?)", *profile.Birthday, key),
	}).Error
}
