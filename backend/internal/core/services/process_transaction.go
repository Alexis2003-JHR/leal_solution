package services

import (
	"context"
	"fmt"
	"leal/internal/core/domain/models/db"
	"leal/internal/core/domain/models/request"
	"time"
)

func (s *service) ProcessTransaction(ctx context.Context, request request.ProcessTransaction) error {
	branch, err := s.repo.FindBranch(ctx, request.BranchID)
	if err != nil {
		return err
	}

	factor, err := s.repo.FindConversionFactor(ctx, branch.BusinessTaxID, request.BranchID)
	if err != nil {
		return err
	}

	now := time.Now()
	campaign, _ := s.repo.FindActiveCampaign(ctx, request.BranchID, now)
	hasCampaign := campaign != nil

	points, cashback := s.calculateEarnings(request.Valor, factor, campaign, hasCampaign)

	tx := &db.Transaction{
		UserDocumentNumber: request.User.DocumentNumber,
		BranchID:           uint(request.BranchID),
		Amount:             request.Valor,
	}

	earnings := &db.Earnings{
		PointsEarned:   points,
		CashbackEarned: cashback,
	}

	if err := s.repo.SaveTransaction(ctx, tx, earnings); err != nil {
		return fmt.Errorf("error saving transaction")
	}

	return nil
}

func (s *service) calculateEarnings(amount float64, factor *db.ConversionFactor, campaign *db.Campaign, hasCampaign bool) (float64, float64) {
	points := amount * factor.PointsPerCurrency
	cashback := amount * factor.CashbackPerCurrency

	if hasCampaign && amount >= campaign.MinPurchaseAmount {
		points *= campaign.PointsMultiplier
		cashback *= campaign.CashbackMultiplier
	}

	return points, cashback
}
