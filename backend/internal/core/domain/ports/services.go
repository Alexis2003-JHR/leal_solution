package ports

import (
	"context"
	"leal/internal/core/domain/models/request"
)

type Service interface {
	CrearUsuario(ctx context.Context, request request.CreateUser) error
}
