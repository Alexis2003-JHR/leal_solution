package api

import (
	"leal/internal/adapters/handlers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, handler *handlers.Handler, txHandler *handlers.TransactionHandler) {
	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/", handler.CreateUser)
			users.GET("/:user/balance", handler.CreateUser)
		}

		campaigns := v1.Group("/campaigns")
		{
			campaigns.POST("/", handler.CreateCampaign)
			campaigns.GET("/", handler.ObtainCampaign)
		}

		branches := v1.Group("/branches")
		{
			branches.POST("/", handler.CreateBranch)
			branches.GET("/", handler.ObtainBranches)
		}

		redemptions := v1.Group("/redemptions")
		{
			redemptions.POST("/points", handler.RedeemPoints)
			redemptions.POST("/cashback", handler.CreateBranch)
		}

		v1.POST("/transactions", txHandler.ProcessTransaction)
		v1.POST("/rewards", handler.CreateReward)
		v1.POST("/business", handler.CreateBusiness)
	}
}
