package utils

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCtxByGin(c *gin.Context) context.Context {
	ctxByGin, exists := c.Get("context")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "context not found"})
		return nil
	}

	ctx, ok := ctxByGin.(context.Context)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid context type"})
		return nil
	}

	return ctx
}
