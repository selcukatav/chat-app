package handler

import (
	"net/http"
	"time"

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
	accessToken, err := middlewares.GenerateToken(&user)
	if err != nil {
		return echo.ErrInternalServerError
	}
	refreshToken, err := middlewares.RefreshToken(&user)
	if err != nil {
		return echo.ErrInternalServerError
	}
	c.Response().Header().Set("Authorization", "Bearer "+accessToken)

	middlewares.SetCookie(c, "access_token", accessToken, time.Now().Add(time.Hour*24))
	middlewares.SetCookie(c, "refresh_token", refreshToken, time.Now().Add(time.Hour*24))

	return c.JSON(http.StatusOK, map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"Message":       "Login successfull!",
	})

}
