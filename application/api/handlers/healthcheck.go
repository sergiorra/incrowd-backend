package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Healthcheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}
