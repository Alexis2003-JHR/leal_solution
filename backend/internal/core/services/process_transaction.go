package services

import (
	"context"
	"fmt"
	"leal/internal/core/domain/models/db"
	"leal/internal/core/domain/models/request"
	"leal/internal/utils"
	"math"
	"time"
)

func (tp *TransactionService) ProcessTransaction(request request.ProcessTransaction) error {
	ctx := context.Background()

	branch, err := tp.repo.FindBranch(ctx, request.BranchID)
	if err != nil {
		return err
	}

	factor, err := tp.repo.FindConversionFactor(ctx, branch.BusinessTaxID, request.BranchID)
	if err != nil {
		return err
	}

	now := time.Now()
	campaign, _ := tp.repo.FindActiveCampaign(ctx, request.BranchID, now)
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
		return fmt.Errorf("error saving transaction")
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
