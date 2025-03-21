package handlers

import (
	"leal/internal/core/domain/custom_errors"
	"leal/internal/core/domain/models"
	"leal/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CrearUsuario(c *gin.Context) {
	ctx := utils.GetCtxByGin(c)
	var requestData models.RequestProof
	if err := c.ShouldBindJSON(&requestData); err != nil {
		log.Println("Error", err)
	}

	result, err := h.Service.CrearUsuario(ctx, requestData)
	if err != nil {
		customErr := custom_errors.New(custom_errors.ErrInternalServerErrorType, "IT Proof", "Something unexpected has happened", http.StatusInternalServerError, err.Error())
		c.Error(customErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": result})
}
