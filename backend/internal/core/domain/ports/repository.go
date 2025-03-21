package ports

import (
	"context"
	"leal/internal/core/domain/models/db"
)

type Repository interface {
	InsertUser(ctx context.Context, user db.User) error
}
