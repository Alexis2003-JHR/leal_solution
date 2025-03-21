package handlers

import (
	"leal/internal/core/domain/models/request"
	"leal/internal/core/services"
	"leal/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionService *services.TransactionService
}

func NewTransactionHandler(transactionService *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
	}
}

func (th *TransactionHandler) ProcessTransaction(c *gin.Context) {
	var tx request.ProcessTransaction
	if err := c.ShouldBindJSON(&tx); err != nil {
		customErr := utils.ParseError(err)
		c.Error(customErr)
		return
	}

	err := th.transactionService.AddTransaction(tx)
	if err != nil {
		customErr := utils.ParseError(err)
		c.Error(customErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transacción encolada"})
}
