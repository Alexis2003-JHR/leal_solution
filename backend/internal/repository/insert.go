package repository

import (
	"context"
	"fmt"
	"leal/internal/core/domain/custom_errors"
	"leal/internal/core/domain/db"
)

func (r *repository) InsertLeal(ctx context.Context, leal db.Leal) error {
	if err := r.db.WithContext(ctx).Create(&leal).Error; err != nil {
		return fmt.Errorf("%w: error inserting the application: %w", custom_errors.ErrSavingError, err)
	}
	return nil
}
