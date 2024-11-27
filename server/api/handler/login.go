package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/selcukatav/chat-app/api/middlewares"
	"github.com/selcukatav/chat-app/models"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) Login(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return echo.ErrBadRequest
	}
	username, userPassword := user.Username, user.UserPassword
	if err := middlewares.Authentication(h.DB, username, userPassword); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "Invalid username or password",
		})

	}
	return c.JSON(http.StatusOK,map[string]string{
		"Message": "Login successfull!",
	})

}
