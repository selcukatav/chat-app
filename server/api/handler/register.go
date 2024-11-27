package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/selcukatav/chat-app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (h *Handler) Register(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return echo.ErrBadRequest
	}
	username, email, userPassword := user.Username, user.Email, user.UserPassword
	result := h.DB.Where("username=? or email=?", username, email).First(&user)
	if result == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "This username is already in use!!",
		})
	} else if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": "Something went wrong",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": "There is error occured while hashing the password",
		})
	}

	user.UserPassword = string(hashedPassword)
	if err := h.DB.Create(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": "Error while registering",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Register successfull!",
	})
}
