package repository

import (
	"context"
	"errors"
	"fmt"
	"leal/internal/core/domain/custom_errors"
	"leal/internal/core/domain/models/db"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r *repository) InsertRedemption(ctx context.Context, redemption db.Redemption) error {
	if err := r.db.WithContext(ctx).Create(&redemption).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return fmt.Errorf("%w: el business_tax_id %d no existe", custom_errors.ErrSavingError, redemption.BusinessTaxID)
		}
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return fmt.Errorf("%w: redemption %d already exists", custom_errors.ErrDuplicatedKey, redemption.UserDocumentNumber)
		}
		return fmt.Errorf("%w: error inserting the redemption", custom_errors.ErrInternalServerError)
	}

	return nil
}
