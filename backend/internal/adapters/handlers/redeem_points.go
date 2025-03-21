package handlers

import (
	"leal/internal/core/domain/models/request"
	"leal/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
