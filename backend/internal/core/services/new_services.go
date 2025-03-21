package services

import (
	"context"
	"leal/internal/adapters/repository"
	"leal/internal/core/domain/models"
)

type Service interface {
	CrearUsuario(ctx context.Context, request models.RequestProof) (int, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}
