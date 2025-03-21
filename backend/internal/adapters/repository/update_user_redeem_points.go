package repository

import (
	"context"
	"leal/internal/core/domain/models/db"

	"gorm.io/gorm"
)

func (r *repository) UpdateUserRedeemPoints(ctx context.Context, userID int, points int) error {
	return r.db.Model(db.UserBalance{}).Where("user_id = ?", userID).UpdateColumn("points", gorm.Expr("points - ?", points)).Error
}
