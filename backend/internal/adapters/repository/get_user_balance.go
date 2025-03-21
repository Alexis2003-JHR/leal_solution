package repository

import (
	"context"
	"fmt"
	"leal/internal/core/domain/custom_errors"
	"leal/internal/core/domain/models/db"
)

func (r *repository) GetUserBalance(ctx context.Context, userID int) (db.UserBalance, error) {
	var userBalance db.UserBalance

	result := r.db.Where("user_id = ?", userID).Find(&userBalance)
	if result.Error != nil {
		return db.UserBalance{}, fmt.Errorf("%w: error al consultar balance: %w", custom_errors.ErrInternalServerError, result.Error)
	}

	return userBalance, nil
}
