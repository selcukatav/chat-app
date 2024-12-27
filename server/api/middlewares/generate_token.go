package middlewares

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/selcukatav/chat-app/model"
)

type jwtCustomClaims struct {
	Username string
	Role     string
	jwt.RegisteredClaims
}

var jwtKey = []byte("very-secret-key")

func GenerateToken(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":      1,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return accessToken, nil

}
func RefreshToken(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":      1,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return refreshToken, nil

}

func SetCookie(c echo.Context, name, value string, expires time.Time) {
	cookie := &http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expires,
		Path:    "/",
	}
	c.SetCookie(cookie)
}
