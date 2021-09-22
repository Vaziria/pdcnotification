package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func SendNotification(c echo.Context) error {
	return c.String(http.StatusOK, "Not Implemented")
}
