package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParamFromContext[T any](context *gin.Context, paramName string) (T, error) {
	paramValue := context.Param(paramName)
	var result T
	switch any(result).(type) {
	case int:
		value, err := strconv.Atoi(paramValue)
		if err != nil {
			return result, err
		}
		return any(value).(T), nil
	case string:
		return any(paramValue).(T), nil
	default:
		return result, fmt.Errorf("unsupported param type")
	}
}
