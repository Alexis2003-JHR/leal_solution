package repository

import (
	"context"
	"fmt"
	"leal/internal/core/domain/custom_errors"
	"leal/internal/core/domain/models/db"
)

func (r *repository) GetReward(ctx context.Context, rewardID int) (db.Reward, error) {
	var reward db.Reward

	result := r.db.Where("id = ?", rewardID).Find(&reward)
	if result.Error != nil {
		return db.Reward{}, fmt.Errorf("%w: error al consultar reward: %w", custom_errors.ErrInternalServerError, result.Error)
	}

	return reward, nil
}
