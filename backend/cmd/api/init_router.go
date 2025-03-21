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
			users.POST("/", handler.AgregarLeal)
			users.GET("/:user/balance", handler.AgregarLeal)
		}

		transactions := v1.Group("/transactions")
		{
			transactions.POST("/", handler.AgregarLeal)
			transactions.GET("/:id", handler.AgregarLeal)
		}

		campaigns := v1.Group("/campaigns")
		{
			campaigns.POST("/campaigns", handler.AgregarLeal)
			campaigns.GET("/campaigns/:comerce_id", handler.AgregarLeal)
		}

		v1.POST("/redemptions", handler.AgregarLeal)
		v1.POST("/merchants", handler.AgregarLeal)
		v1.POST("/branches", handler.AgregarLeal)
	}
}
