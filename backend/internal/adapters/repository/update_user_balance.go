package repository

import "context"

func (r *repository) UpdateUserBalance(ctx context.Context, userID int, points int, cashback float64) error {
	return r.db.WithContext(ctx).
		Exec(`
			INSERT INTO user_balances (user_id, points, cashback, updated_at)
			VALUES ($1, $2, $3, NOW())
			ON CONFLICT (user_id) 
			DO UPDATE SET 
				points = user_balances.points + EXCLUDED.points,
				cashback = user_balances.cashback + EXCLUDED.cashback,
				updated_at = NOW()
		`, userID, points, cashback).Error
}
