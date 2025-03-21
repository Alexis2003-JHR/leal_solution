package repository

import (
	"context"
	"fmt"
	"leal/internal/core/domain/custom_errors"
	"leal/internal/core/domain/models/db"
)

func (r *repository) GetCampaigns(ctx context.Context, taxID int) ([]db.Campaign, error) {
	var campaigns []db.Campaign

	result := r.db.Where("business_tax_id = ?", taxID).Find(&campaigns)
	if result.Error != nil {
		return nil, fmt.Errorf("%w: error al consultar campañas: %w", custom_errors.ErrInternalServerError, result.Error)
	}

	if len(campaigns) == 0 {
		return nil, fmt.Errorf("%w: no se encontraron campañas para el tax_id: %d", custom_errors.ErrNotFound, taxID)
	}

	return campaigns, nil
}
