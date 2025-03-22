package handlers

import (
	"leal/internal/core/domain/models/request"
	"leal/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RedeemPoints		godoc
// @Summary			Redimir puntos
// @Description		Permitir a un usuario redimir puntos por un premio específico.
// @Param			redeem body request.RedeemPoints true "Datos para redimir puntos"
// @Produce			application/json
// @Tags			Rewards
// @Success			200
// @Router			/redemptions/points [post]
func (h *Handler) RedeemPoints(c *gin.Context) {
	ctx := utils.GetCtxByGin(c)
	var requestData request.RedeemPoints
	if err := c.ShouldBindJSON(&requestData); err != nil {
		customErr := utils.ParseError(err)
		c.Error(customErr)
		return
	}

	points, err := h.Service.RedeemPoints(ctx, requestData)
	if err != nil {
		customErr := utils.ParseError(err)
		c.Error(customErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Premio redimido correctamente", "puntos_redimidos": points})
}
