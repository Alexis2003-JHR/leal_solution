package repository

import (
	"context"
	"errors"
	"fmt"
	"leal/internal/core/domain/custom_errors"
	"leal/internal/core/domain/models/db"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

func (r *repository) InsertCampaign(ctx context.Context, campaign db.Campaign) error {
	var existing db.Campaign
	err := r.db.WithContext(ctx).
		Where("business_tax_id = ? AND branch_id = ? AND NOT (end_date < ? OR start_date > ?)",
			campaign.BusinessTaxID, campaign.BranchID, campaign.StartDate, campaign.EndDate).
		First(&existing).Error

	if err == nil {
		return fmt.Errorf("%w: ya existe una campaña en este rango de fechas para la sucursal %d del negocio %d",
			custom_errors.ErrDuplicatedKey, campaign.BranchID, campaign.BusinessTaxID)
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("%w: error al validar la existencia de campañas: %w", custom_errors.ErrInternalServerError, err)
	}

	if err := r.db.WithContext(ctx).Create(&campaign).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return fmt.Errorf("%w: el business_tax_id %d no existe", custom_errors.ErrSavingError, campaign.BusinessTaxID)
		}
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return fmt.Errorf("%w: campaign to %d already exists", custom_errors.ErrDuplicatedKey, campaign.BusinessTaxID)
		}
		return fmt.Errorf("%w: error inserting the campaign", custom_errors.ErrInternalServerError)
	}

	return nil
}
