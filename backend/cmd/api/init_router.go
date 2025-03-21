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
			users.POST("/", handler.CreateUser)
			users.GET("/:user/balance", handler.CreateUser)
		}

		transactions := v1.Group("/transactions")
		{
			transactions.POST("/", handler.CreateUser)
			transactions.GET("/:id", handler.CreateUser)
		}

		campaigns := v1.Group("/campaigns")
		{
			campaigns.POST("/campaigns", handler.CreateCampaign)
			campaigns.GET("/campaigns/:comerce_id", handler.CreateUser)
		}

		v1.POST("/redemptions", handler.CreateUser)
		v1.POST("/business", handler.CreateBusiness)
		v1.POST("/branches", handler.CreateUser)
	}
}
