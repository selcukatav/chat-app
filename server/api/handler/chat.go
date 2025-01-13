package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Rooms(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Welcome to the restricted page",
	})
}
