package handlers

import "leal/internal/core/services"

type Handler struct {
	Service services.Service
}

func NewHandler(service services.Service) *Handler {
	return &Handler{Service: service}
}
