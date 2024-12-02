package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/selcukatav/chat-app/api/middlewares"
	"github.com/selcukatav/chat-app/model"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) Login(c echo.Context) error {
	var user model.User
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
