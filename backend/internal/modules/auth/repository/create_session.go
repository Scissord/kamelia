package repository

import (
	"backend/internal/modules/session/model"
	"context"
)

func (r *repository) CreateSession(ctx context.Context, s *model.Session) error {
	return r.db.WithContext(ctx).Create(s).Error
}
