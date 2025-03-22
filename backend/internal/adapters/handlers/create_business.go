package handlers

import (
	"leal/internal/core/domain/models/request"
	"leal/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateBusiness godoc
// @Summary      Crear nuevo comercio
// @Description  Endpoint to register a new business. Requires business details in the request body.
// @Tags         Businesses
// @Accept       json
// @Produce      json
// @Param        business body request.CreateBusiness true "Business data"
// @Success      200
// @Router       /business [post]
func (h *Handler) CreateBusiness(c *gin.Context) {
	ctx := utils.GetCtxByGin(c)
	var requestData request.CreateBusiness
	if err := c.ShouldBindJSON(&requestData); err != nil {
		customErr := utils.ParseError(err)
		c.Error(customErr)
		return
	}

	err := h.Service.CreateBusiness(ctx, requestData)
	if err != nil {
		customErr := utils.ParseError(err)
		c.Error(customErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Empresa creada correctamente"})
}
