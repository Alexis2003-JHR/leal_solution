package ports

import (
	"context"
	"leal/internal/core/domain/models/db"
)

type Repository interface {
	InsertUser(ctx context.Context, user db.User) error
	InsertBusiness(ctx context.Context, business db.Business) error
	InsertBranch(ctx context.Context, business db.Branch) error
	GetBranches(ctx context.Context, taxID int) ([]db.Branch, error)
}
