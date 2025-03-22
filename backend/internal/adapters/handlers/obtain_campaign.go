package handlers

import (
	"leal/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ObtainCampaign	godoc
// @Summary			Obtener campañas de comercio
// @Description		Recuperar las campañas asociadas a una empresa específica.
// @Param			tax_id query int true "ID tributario de la empresa"
// @Produce			application/json
// @Tags			Campaigns
// @Success			200
// @Router			/campaigns [get]
func (h *Handler) ObtainCampaign(c *gin.Context) {
	ctx := utils.GetCtxByGin(c)
	idParam := c.Query("tax_id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tax_id es requerido"})
		return
	}
	taxID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	result, err := h.Service.ObtainCampaign(ctx, taxID)
	if err != nil {
		customErr := utils.ParseError(err)
		c.Error(customErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"campañas": result})
}
