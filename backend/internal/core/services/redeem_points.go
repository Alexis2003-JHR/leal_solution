package services

import (
	"context"
	"fmt"
	"leal/internal/core/domain/models/db"
	"leal/internal/core/domain/models/request"
)

func (s *service) RedeemPoints(ctx context.Context, request request.RedeemPoints) (int, error) {
	userBalance, err := s.repo.GetUserBalance(ctx, request.User.DocumentNumber)
	if err != nil {
		return 0, err
	}

	reward, err := s.repo.GetReward(ctx, request.RewardID)
	if err != nil {
		return 0, err
	}

	if userBalance.Points < reward.PointsRequired {
		return 0, fmt.Errorf("puntos insuficientes para reclamar el premio")
	}

	if err := s.repo.UpdateUserRedeemPoints(ctx, request.User.DocumentNumber, reward.PointsRequired); err != nil {
		return 0, err
	}

	redemption := db.Redemption{
		UserDocumentNumber: request.User.DocumentNumber,
		BusinessTaxID:      request.BusinessTaxID,
		RewardID:           uint(request.RewardID),
		PointsSpent:        reward.PointsRequired,
	}

	if err := s.repo.InsertRedemption(ctx, redemption); err != nil {
		return 0, err
	}

	return reward.PointsRequired, nil
}
