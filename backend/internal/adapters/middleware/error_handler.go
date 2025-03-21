package middleware

import (
	"leal/internal/core/domain/custom_errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			if customErr, ok := err.(*custom_errors.LogError); ok {
				logError(customErr)
				appErr := custom_errors.AppError{
					Code:    string(customErr.Type),
					Message: customErr.Message,
				}
				c.JSON(customErr.StatusCode, gin.H{
					"error": appErr,
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
	}
}

func logError(err *custom_errors.LogError) {
	log.Printf(`{"level": "error", "type": "%s", "message": "%s", "log_message": "%s"}`, err.Type, err.Message, err.LogMessage)
}
