package services

import (
	"context"
	"leal/internal/core/domain/models/db"
	"leal/internal/core/domain/models/request"
)

func (s *service) CreateBusiness(ctx context.Context, request request.CreateBusiness) error {
	business := db.Business{
		Name:  request.RazonSocial,
		TaxID: request.NIT,
		Phone: request.Telefono,
		Email: request.Correo,
	}

	err := s.repo.InsertBusiness(ctx, business)
	if err != nil {
		return err
	}

	conversionFactor := db.ConversionFactor{
		BusinessTaxID:       request.NIT,
		MinAmount:           request.ConversionFactor.MinAmount,
		PointsPerCurrency:   request.ConversionFactor.PointsPerCurrency,
		CashbackPerCurrency: request.ConversionFactor.CashbackPerCurrency,
	}

	err = s.repo.InsertConversionFactor(ctx, conversionFactor)
	if err != nil {
		return err
	}

	return nil
}
