package repository

import (
	"context"
	"errors"
	"fmt"
	"leal/internal/core/domain/custom_errors"
	"leal/internal/core/domain/models/db"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r *repository) InsertUser(ctx context.Context, user db.User) error {
	if user.Email == "" || user.Name == "" {
		return custom_errors.ErrBadRequest
	}

	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return fmt.Errorf("%w: email %s already exists", custom_errors.ErrDuplicatedKey, user.Email)
		}
		return fmt.Errorf("%w: error inserting user", custom_errors.ErrInternalServerError)
	}

	return nil
}
