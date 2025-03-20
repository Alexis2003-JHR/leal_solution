package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ParseRequestWithTurno[T any](c *gin.Context) (*T, error) {
	var requestData T

	turnoId, exists := c.Get("turno")
	if !exists {
		return nil, fmt.Errorf("turnoId not found in context")
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		return nil, err
	}

	if v, ok := any(&requestData).(interface{ SetTurnoId(int) }); ok {
		v.SetTurnoId(turnoId.(int))
	}

	return &requestData, nil
}
