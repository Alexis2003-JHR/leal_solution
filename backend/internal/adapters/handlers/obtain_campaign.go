package handlers

import (
	"fmt"
	"leal/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ObtainCampaign	godoc
// @Summary			Obtener campañas de comercio
// @Description		Recuperar las campañas asociadas a una empresa específica.
// @Param			tax_id query int true "ID tributario de la empresa"
// @Param			branch_id query int false "ID sucursal"
// @Produce			application/json
// @Tags			Campaigns
// @Success			200
// @Router			/campaigns [get]
func (h *Handler) ObtainCampaign(c *gin.Context) {
	ctx := utils.GetCtxByGin(c)

	taxID, branchID, err := parseCampaignRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.Service.ObtainCampaign(ctx, taxID, branchID)
	if err != nil {
		customErr := utils.ParseError(err)
		c.Error(customErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"campañas": result})
}

func parseCampaignRequest(c *gin.Context) (int, *int, error) {
	idParam := c.Query("tax_id")
	if idParam == "" {
		return 0, nil, fmt.Errorf("tax_id es requerido")
	}

	taxID, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, nil, fmt.Errorf("tax_id inválido")
	}

	branchParam := c.Query("branch_id")
	var branchID *int
	if branchParam != "" {
		parsedBranchID, err := strconv.Atoi(branchParam)
		if err != nil {
			return 0, nil, fmt.Errorf("branch_id inválido")
		}
		branchID = &parsedBranchID
	}

	return taxID, branchID, nil
}
