package services

import (
	"context"
	"leal/internal/core/domain/models"
)

func (s *service) ObtainCampaign(ctx context.Context, taxID int) ([]models.Campaign, error) {
	campaignsDB, err := s.repo.GetCampaigns(ctx, taxID)
	if err != nil {
		return []models.Campaign{}, err
	}

	var campaigns []models.Campaign
	for _, campaignDB := range campaignsDB {
		campaign := models.Campaign{
			ID:                 int(campaignDB.ID),
			BranchID:           campaignDB.BusinessTaxID,
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
