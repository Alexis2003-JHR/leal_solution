package services

import (
	"context"
	"leal/internal/core/domain/models/db"
	"leal/internal/core/domain/models/request"
)

func (s *service) CreateUser(ctx context.Context, request request.CreateUser) error {
	user := db.User{
		Name:           request.Name,
		DocumentNumber: request.DocumentNumber,
		Email:          request.Email,
	}
	return s.repo.InsertUser(ctx, user)
}
