package handlers

import (
	"leal/internal/core/domain/models/request"
	"leal/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCampaign godoc
// @Summary      Crear nueva campaña
// @Description  Endpoint to create a new marketing campaign for a business. Requires campaign details in the request body.
// @Tags         Campaigns
// @Accept       json
// @Produce      json
// @Param        campaign body request.CreateCampaign true "Campaign data"
// @Success      200
// @Router       /campaigns [post]
func (h *Handler) CreateCampaign(c *gin.Context) {
	ctx := utils.GetCtxByGin(c)
	var requestData request.CreateCampaign
	if err := c.ShouldBindJSON(&requestData); err != nil {
		customErr := utils.ParseError(err)
		c.Error(customErr)
		return
	}

	err := h.Service.CreateCampaign(ctx, requestData)
	if err != nil {
		customErr := utils.ParseError(err)
		c.Error(customErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Campaña ha sido creada correctamente"})
}
