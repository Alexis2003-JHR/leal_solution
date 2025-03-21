package ports

import (
	"context"
	"leal/internal/core/domain/models/db"
)

type Repository interface {
	InsertConversionFactor(ctx context.Context, business db.ConversionFactor) error
	InsertBusiness(ctx context.Context, business db.Business) error
	InsertBranch(ctx context.Context, business *db.Branch) error
	InsertUser(ctx context.Context, user db.User) error
	GetBranches(ctx context.Context, taxID int) ([]db.Branch, error)
	InsertCampaign(ctx context.Context, campaign db.Campaign) error
}
