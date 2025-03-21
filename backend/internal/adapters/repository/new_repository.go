package repository

import (
	"leal/internal/core/domain/ports"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewPostgrestRepository(db *gorm.DB) ports.Repository {
	return &repository{db: db}
}
