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
	return s.repo.InsertBusiness(ctx, business)
}
