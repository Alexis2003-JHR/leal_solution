package api

import (
	"leal/internal/adapters/handlers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, handler *handlers.Handler) {
	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/", handler.CrearUsuario)
			users.GET("/:user/balance", handler.CrearUsuario)
		}

		transactions := v1.Group("/transactions")
		{
			transactions.POST("/", handler.CrearUsuario)
			transactions.GET("/:id", handler.CrearUsuario)
		}

		campaigns := v1.Group("/campaigns")
		{
			campaigns.POST("/campaigns", handler.CrearUsuario)
			campaigns.GET("/campaigns/:comerce_id", handler.CrearUsuario)
		}

		v1.POST("/redemptions", handler.CrearUsuario)
		v1.POST("/merchants", handler.CrearUsuario)
		v1.POST("/branches", handler.CrearUsuario)
	}
}
