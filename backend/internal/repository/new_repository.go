package repository

import (
	"context"
	"leal/internal/core/domain/db"

	"gorm.io/gorm"
)

type Repository interface {
	InsertLeal(ctx context.Context, leal db.Leal) error
}

type repository struct {
	db *gorm.DB
}

func NewPostgrestRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
