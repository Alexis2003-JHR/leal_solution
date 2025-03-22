package handlers

import (
	"leal/internal/core/domain/models/request"
	"leal/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateBranch godoc
// @Summary      Crear nueva sucursal
// @Description  Endpoint to create a new branch for a business. Requires branch details in the request body.
// @Tags         Branches
// @Accept       json
// @Produce      json
// @Param        branch body request.CreateBranch true "Branch data"
// @Success      200
// @Router       /branches [post]
func (h *Handler) CreateBranch(c *gin.Context) {
	ctx := utils.GetCtxByGin(c)
	var requestData request.CreateBranch
	if err := c.ShouldBindJSON(&requestData); err != nil {
		customErr := utils.ParseError(err)
		c.Error(customErr)
		return
	}

	err := h.Service.CreateBranch(ctx, requestData)
	if err != nil {
		customErr := utils.ParseError(err)
		c.Error(customErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sucursal creada correctamente"})
}
