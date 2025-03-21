package services

import (
	"leal/internal/core/domain/ports"
)

type service struct {
	repo ports.Repository
}

func NewService(repo ports.Repository) ports.Service {
	return &service{repo: repo}
}
