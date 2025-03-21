package repository

import (
	"context"
	"fmt"
	"leal/internal/core/domain/custom_errors"
	"leal/internal/core/domain/models/db"
)

func (r *repository) InsertUser(ctx context.Context, user db.User) error {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return fmt.Errorf("%w: error inserting the application: %w", custom_errors.ErrSavingError, err)
	}
	return nil
}
