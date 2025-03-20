package services

import (
	"context"
	"leal/internal/core/domain"
	"leal/internal/repository"
)

type Service interface {
	AgregarLeal(ctx context.Context, request domain.RequestProof) (int, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}
