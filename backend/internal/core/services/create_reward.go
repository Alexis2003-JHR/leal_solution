package services

import (
	"context"
	"leal/internal/core/domain/models/db"
	"leal/internal/core/domain/models/request"
)

func (s *service) CreateReward(ctx context.Context, request request.CreateReward) error {
	rewardDB := db.Reward{
		Name:           request.Name,
		Description:    request.Description,
		PointsRequired: request.PointsRequired,
		BusinessTaxID:  int64(request.BusinessTaxID),
	}
	return s.repo.InsertReward(ctx, &rewardDB)
}
