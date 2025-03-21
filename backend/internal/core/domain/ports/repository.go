package ports

import (
	"context"
	"leal/internal/core/domain/models/db"
)

type Repository interface {
	InsertConversionFactor(ctx context.Context, business db.ConversionFactor) error
	InsertCampaign(ctx context.Context, campaign db.Campaign) error
	InsertBusiness(ctx context.Context, business db.Business) error
	InsertBranch(ctx context.Context, business *db.Branch) error
	InsertUser(ctx context.Context, user db.User) error
	GetBranches(ctx context.Context, taxID int) ([]db.Branch, error)
	GetCampaigns(ctx context.Context, taxID int) ([]db.Campaign, error)
	FindBranch(ctx context.Context, branchID int) (*db.Branch, error)
}
