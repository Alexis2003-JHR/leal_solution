package handlers

import (
	"leal/internal/core/domain/models/request"
	"leal/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser		godoc
// @Summary			Crear nuevo usuario
// @Description		Guardar los datos de un nuevo usuario en la base de datos.
// @Param			user body request.CreateUser true "Datos del usuario"
// @Produce			application/json
// @Tags			Users
// @Success			200
// @Router			/users [post]
func (h *Handler) CreateUser(c *gin.Context) {
	ctx := utils.GetCtxByGin(c)
	var requestData request.CreateUser
	if err := c.ShouldBindJSON(&requestData); err != nil {
		customErr := utils.ParseError(err)
		c.Error(customErr)
		return
	}

	err := h.Service.CreateUser(ctx, requestData)
	if err != nil {
		customErr := utils.ParseError(err)
		c.Error(customErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario creado correctamente"})
}
