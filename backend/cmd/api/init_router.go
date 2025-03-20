package api

import (
	"leal/internal/core/handlers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, handler *handlers.Handler) {
	v1 := r.Group("/api/v1")
	{
		grupos := v1.Group("/campaña")
		{
			grupos.GET("/", handler.AgregarLeal)
		}
	}
}
