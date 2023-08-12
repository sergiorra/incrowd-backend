package apiutils

import (
	"encoding/json"
	"incrowd-backend/internal/common"

	"github.com/labstack/echo/v4"
)

// ReadStringQueryParam reads string query params. If they are bad formatted, default value is returned
func ReadStringQueryParam(c echo.Context, key, defaultValue string, allowedValues []string) string {
	val := c.QueryParam(key)
	if val == "" || val == "id" {
		return defaultValue
	}

	if allowedValues == nil {
		return val
	}

	if !common.Contains(allowedValues, val) {
		return defaultValue
	}

	return val
}

// ReadIntQueryParam reads int query params. If they are bad formatted, default value is returned
func ReadIntQueryParam(c echo.Context, key string, defaultValue int) int {
	val := c.QueryParam(key)
	if val == "" {
		return defaultValue
	}

	var intVal int
	if err := json.Unmarshal([]byte(val), &intVal); err != nil {
		return defaultValue
	}

	if intVal < 0 {
		return defaultValue
	}

	return intVal
}
