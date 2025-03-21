package ports

import (
	"context"
	"leal/internal/core/domain/models"
	"leal/internal/core/domain/models/request"
)

type Service interface {
	CreateUser(ctx context.Context, request request.CreateUser) error
	CreateCampaign(ctx context.Context, request request.CreateCampaign) error
	CreateBusiness(ctx context.Context, request request.CreateBusiness) error
	CreateBranch(ctx context.Context, request request.CreateBranch) error
	ObtainBranches(ctx context.Context, taxID int) ([]models.Branch, error)
	ObtainCampaign(ctx context.Context, taxID int) ([]models.Campaign, error)
	ProcessTransaction(ctx context.Context, request request.ProcessTransaction) error
}
