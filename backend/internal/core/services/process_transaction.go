package services

import (
	"context"
	"leal/internal/core/domain/models/db"
	"leal/internal/core/domain/models/request"
	"leal/internal/utils"
	"math"
	"time"
)

func (tp *TransactionService) ProcessTransaction(request request.ProcessTransaction) error {
	ctx := context.Background()

	factor, err := tp.repo.FindConversionFactor(ctx, request.BusinessTaxID, request.BranchID)
	if err != nil {
		return err
	}

	now := time.Now()
	campaign, _ := tp.repo.FindActiveCampaign(ctx, request.BusinessTaxID, request.BranchID, now)
	hasCampaign := campaign != nil

	points, cashback := tp.calculateEarnings(request.Valor, factor, campaign, hasCampaign)

	tx := &db.Transaction{
		UserDocumentNumber: request.User.DocumentNumber,
		BranchID:           uint(request.BranchID),
		Amount:             request.Valor,
	}

	earnings := &db.Earnings{
		PointsEarned:   points,
		CashbackEarned: utils.RoundToTwoDecimals(cashback),
	}

	if err := tp.repo.SaveTransaction(ctx, tx, earnings); err != nil {
		return err
	}

	if err := tp.repo.UpdateUserBalance(ctx, request.User.DocumentNumber, earnings.PointsEarned, earnings.CashbackEarned); err != nil {
		return err
	}

	return nil
}

func (s *TransactionService) calculateEarnings(amount float64, factor *db.ConversionFactor, campaign *db.Campaign, hasCampaign bool) (int, float64) {
	points := amount * factor.PointsPerCurrency
	cashback := amount * factor.CashbackPerCurrency

	if hasCampaign && amount >= campaign.MinPurchaseAmount {
		points *= campaign.PointsMultiplier
		cashback *= campaign.CashbackMultiplier
	}

	return int(math.Round(points)), cashback
}
