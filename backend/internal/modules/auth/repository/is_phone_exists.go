package repository

import (
	profileModel "backend/internal/modules/profile/model"
	"backend/internal/utils"
	"context"
)

func (r *repository) IsPhoneExists(ctx context.Context, phone string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&profileModel.Profile{}).
		Where("phone_hash = ?", utils.Hash(&phone)).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
