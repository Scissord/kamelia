package repository

import (
	profileModel "backend/internal/modules/profile/model"
	"backend/internal/utils"
	"context"
)

func (r *repository) IsEmailExists(ctx context.Context, email string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&profileModel.Profile{}).
		Where("email_hash = ?", utils.Hash(&email)).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
