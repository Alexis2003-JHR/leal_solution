package services

import (
	"context"
	"leal/internal/core/domain/models"
)

func (s *service) ObtainCampaign(ctx context.Context, taxID int, branchID *int) ([]models.Campaign, error) {
	campaignsDB, err := s.repo.GetCampaigns(ctx, taxID, branchID)
	if err != nil {
		return []models.Campaign{}, err
	}

	var campaigns []models.Campaign
	for _, campaignDB := range campaignsDB {
		var branchID *int
		if campaignDB.BranchID != nil {
			bID := int(*campaignDB.BranchID)
			branchID = &bID
		}

		campaign := models.Campaign{
			ID:                 int(campaignDB.ID),
			BranchID:           branchID,
			StartDate:          campaignDB.StartDate.Format("2006-01-02"),
			EndDate:            campaignDB.EndDate.Format("2006-01-02"),
			PointsMultiplier:   campaignDB.PointsMultiplier,
			CashbackMultiplier: campaignDB.CashbackMultiplier,
			MinPurchaseAmount:  campaignDB.MinPurchaseAmount,
		}
		campaigns = append(campaigns, campaign)
	}

	return campaigns, nil
}
