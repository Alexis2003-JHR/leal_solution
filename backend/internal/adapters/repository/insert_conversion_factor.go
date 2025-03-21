package repository

import (
	"context"
	"errors"
	"fmt"
	"leal/internal/core/domain/custom_errors"
	"leal/internal/core/domain/models/db"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r *repository) InsertConversionFactor(ctx context.Context, conversionFactor db.ConversionFactor) error {
	if err := r.db.WithContext(ctx).Create(&conversionFactor).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return fmt.Errorf("%w: conversion factor to %d already exists", custom_errors.ErrDuplicatedKey, conversionFactor.BusinessTaxID)
		}
		return fmt.Errorf("%w: error inserting the conversion factor", custom_errors.ErrInternalServerError)
	}

	return nil
}
