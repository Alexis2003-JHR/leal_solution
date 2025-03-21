package ports

import (
	"context"
	"leal/internal/core/domain/models/db"
	"time"
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
	FindConversionFactor(ctx context.Context, businessTaxID int, branchID int) (*db.ConversionFactor, error)
	FindActiveCampaign(ctx context.Context, branchID int, now time.Time) (*db.Campaign, error)
	SaveTransaction(ctx context.Context, tx *db.Transaction, earnings *db.Earnings) error
	UpdateUserBalance(ctx context.Context, userID int, points int, cashback float64) error
}
