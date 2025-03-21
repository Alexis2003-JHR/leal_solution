package handlers

import (
	"leal/internal/core/domain/ports"
)

type Handler struct {
	Service ports.Service
}

func NewHandler(service ports.Service) *Handler {
	return &Handler{Service: service}
}
