package repository

import (
	"context"
	"fmt"
	"leal/internal/core/domain/custom_errors"
	"leal/internal/core/domain/models/db"
)

func (r *repository) GetCampaigns(ctx context.Context, taxID int, branchID *int) ([]db.Campaign, error) {
	var campaigns []db.Campaign

	query := r.db.Where("business_tax_id = ?", taxID)
	if branchID != nil {
		query = query.Where("branch_id = ?", *branchID)
	}

	result := query.Find(&campaigns)
	if result.Error != nil {
		return nil, fmt.Errorf("%w: error al consultar campañas: %w", custom_errors.ErrInternalServerError, result.Error)
	}

	if len(campaigns) == 0 {
		return nil, fmt.Errorf("%w: no se encontraron campañas para el tax_id: %d", custom_errors.ErrNotFound, taxID)
	}

	return campaigns, nil
}
