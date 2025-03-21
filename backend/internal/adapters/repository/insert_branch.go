package repository

import (
	"context"
	"errors"
	"fmt"
	"leal/internal/core/domain/custom_errors"
	"leal/internal/core/domain/models/db"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r *repository) InsertBranch(ctx context.Context, branch db.Branch) error {
	if branch.Name == "" {
		return custom_errors.ErrBadRequest
	}

	if err := r.db.WithContext(ctx).Create(&branch).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return fmt.Errorf("%w: branch %s already exists", custom_errors.ErrDuplicatedKey, branch.Name)
		}
		return fmt.Errorf("%w: error inserting the branch", custom_errors.ErrInternalServerError)
	}

	return nil
}
