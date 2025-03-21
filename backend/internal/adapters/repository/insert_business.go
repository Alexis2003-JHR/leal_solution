package repository

import (
	"context"
	"errors"
	"fmt"
	"leal/internal/core/domain/custom_errors"
	"leal/internal/core/domain/models/db"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r *repository) InsertBusiness(ctx context.Context, business db.Business) error {
	if business.Name == "" {
		return custom_errors.ErrBadRequest
	}

	var existing db.Business
	if err := r.db.WithContext(ctx).Where("name = ? OR tax_id = ? OR email = ?", business.Name, business.TaxID, business.Email).First(&existing).Error; err == nil {
		return fmt.Errorf("%w: business with the same name, tax ID, or email already exists", custom_errors.ErrDuplicatedKey)
	}

	if err := r.db.WithContext(ctx).Create(&business).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return fmt.Errorf("%w: business %s already exists", custom_errors.ErrDuplicatedKey, business.Name)
		}
		return fmt.Errorf("%w: error inserting the business", custom_errors.ErrInternalServerError)
	}

	return nil
}
