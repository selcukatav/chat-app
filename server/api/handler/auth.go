package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/selcukatav/chat-app/api/middlewares"
	"github.com/selcukatav/chat-app/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

// @Summary      User login
// @Description  User logs in and gets token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        username  body  string  true  "username"
// @Param        userPassword  body  string  true  "userPassword"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Router       /api/login [get]
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

// @Summary      User Register
// @Description  User Registers and gets token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        username  body  string  true  "username"
// @Param        email  body  string  true  "email"
// @Param        userPassword  body  string  true  "userPassword"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Router       /api/register [post]
func (h *Handler) Register(c echo.Context) error {
	var user model.User

	if err := c.Bind(&user); err != nil {
		return echo.ErrBadRequest
	}
	username, email, userPassword := user.Username, user.Email, user.UserPassword
	result := h.DB.Where("username=? or email=?", username, email).First(&user)
	if result.RowsAffected > 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "This username or email is already in use!!",
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
	createUser := h.DB.Create(&user)

	if createUser.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": "There is error occured while creating user",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{

		"Message": "Register successfull!",
	})
}
