package repository

import (
	"context"
	"fmt"
	"leal/internal/core/domain/custom_errors"
	"leal/internal/core/domain/models/db"
)

func (r *repository) GetBranches(ctx context.Context, taxID int) ([]db.Branch, error) {
	var branches []db.Branch

	result := r.db.Where("business_tax_id = ?", taxID).Find(&branches)
	if result.Error != nil {
		return nil, fmt.Errorf("%w: error al consultar branches: %w", custom_errors.ErrInternalServerError, result.Error)
	}

	if len(branches) == 0 {
		return nil, fmt.Errorf("%w: no se encontraron sucursales para el tax_id: %d", custom_errors.ErrNotFound, taxID)
	}

	return branches, nil
}
