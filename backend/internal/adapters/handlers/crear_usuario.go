package handlers

import (
	"leal/internal/core/domain/custom_errors"
	"leal/internal/core/domain/models/request"
	"leal/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CrearUsuario(c *gin.Context) {
	ctx := utils.GetCtxByGin(c)
	var requestData request.CreateUser
	if err := c.ShouldBindJSON(&requestData); err != nil {
		customErr := utils.ParseError(err)
		c.Error(customErr)
		return
	}

	err := h.Service.CrearUsuario(ctx, requestData)
	if err != nil {
		customErr := custom_errors.New(custom_errors.ErrInternalServerErrorType, "Something unexpected has happened", http.StatusInternalServerError, err.Error())
		c.Error(customErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario creado correctamente"})
}
