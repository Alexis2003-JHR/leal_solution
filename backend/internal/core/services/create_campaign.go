package services

import (
	"context"
	"fmt"
	"leal/internal/core/domain/custom_errors"
	"leal/internal/core/domain/models/db"
	"leal/internal/core/domain/models/request"
	"time"
)

func (s *service) CreateCampaign(ctx context.Context, request request.CreateCampaign) error {
	parsedStartDate, err := time.Parse("2006-01-02", request.StartDate)
	if err != nil {
		return fmt.Errorf("%w: invalid date format: %w", custom_errors.ErrBadRequest, err)
	}

	parsedEndDate, err := time.Parse("2006-01-02", request.EndDate)
	if err != nil {
		return fmt.Errorf("%w: invalid date format: %w", custom_errors.ErrBadRequest, err)
	}

	var branchID *uint
	if request.BranchID > 0 {
		tempBranchID := uint(request.BranchID)
		branchID = &tempBranchID
	}

	campaignDB := db.Campaign{
		BusinessTaxID:      request.TaxID,
		BranchID:           branchID,
		StartDate:          parsedStartDate,
		EndDate:            parsedEndDate,
		PointsMultiplier:   request.PointsMultiplier,
		CashbackMultiplier: request.CashbackMultiplier,
		MinPurchaseAmount:  request.MinPurchaseAmount,
	}

	err = s.repo.InsertCampaign(ctx, campaignDB)
	if err != nil {
		return err
	}

	return nil
}
