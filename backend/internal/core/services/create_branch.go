package services

import (
	"context"
	"leal/internal/core/domain/models/db"
	"leal/internal/core/domain/models/request"
)

func (s *service) CreateBranch(ctx context.Context, request request.CreateBranch) error {
	branch := db.Branch{
		BusinessTaxID: request.NITEmpresa,
		Name:          request.NombreSucursal,
	}

	err := s.repo.InsertBranch(ctx, &branch)
	if err != nil {
		return err
	}

	conversionFactor := db.ConversionFactor{
		BusinessTaxID:       request.NITEmpresa,
		BranchID:            &branch.ID,
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
