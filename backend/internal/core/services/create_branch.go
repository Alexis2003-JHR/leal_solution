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
	return s.repo.InsertBranch(ctx, branch)
}
