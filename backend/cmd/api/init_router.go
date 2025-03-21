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
			campaigns.POST("/", handler.CreateCampaign)
			campaigns.GET("/", handler.CreateUser)
		}

		branches := v1.Group("/branches")
		{
			branches.POST("/", handler.CreateBranch)
			branches.GET("/", handler.ObtainBranches)
		}

		v1.POST("/redemptions", handler.CreateUser)
		v1.POST("/business", handler.CreateBusiness)
	}
}
