package handlers

import (
	"leal/internal/core/domain/models/request"
	"leal/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateReward		godoc
// @Summary			Crear nuevo premio
// @Description		Guardar los datos de un nuevo premio en la base de datos.
// @Param			reward body request.CreateReward true "Datos del premio"
// @Produce			application/json
// @Tags			Rewards
// @Success			200
// @Router			/rewards [post]
func (h *Handler) CreateReward(c *gin.Context) {
	ctx := utils.GetCtxByGin(c)
	var requestData request.CreateReward
	if err := c.ShouldBindJSON(&requestData); err != nil {
		customErr := utils.ParseError(err)
		c.Error(customErr)
		return
	}

	err := h.Service.CreateReward(ctx, requestData)
	if err != nil {
		customErr := utils.ParseError(err)
		c.Error(customErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Premio creado correctamente"})
}
