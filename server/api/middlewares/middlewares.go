package middlewares

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/selcukatav/chat-app/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Authentication(postgres *gorm.DB, username, password string) error {

	var user model.User
	result := postgres.Where("username=?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(password))
	if err != nil {
		return errors.New("wrong password")
	}

	return nil
}

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

func Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := authHeader(c); err == nil {
			return next(c)
		}
		if err := authCookie(c); err == nil {
			return next(c)
		}
		return echo.ErrUnauthorized
	}
}

func authHeader(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	claims := jwt.MapClaims{}
	if authHeader != "" {
		tokenHeader := authHeader[len("Bearer "):]

		token, err := jwt.ParseWithClaims(tokenHeader, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.ErrUnauthorized
			}
			return jwtKey, nil
		})

		if err == nil && token.Valid {
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				sub := int(claims["sub"].(float64))
				if sub != 1 {
					return echo.ErrForbidden
				}
				return nil
			}
		}
	}

	return echo.ErrUnauthorized
}

func authCookie(c echo.Context) error {
	authCookie, err := c.Cookie("access_token")
	if err != nil || authCookie == nil {
		return echo.ErrUnauthorized
	}
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(authCookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.ErrUnauthorized
		}
		return jwtKey, nil
	})

	if err == nil && token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			sub := int(claims["sub"].(float64))
			if sub != 1 {
				return echo.ErrForbidden
			}
			return nil
		}
	}

	return echo.ErrUnauthorized
}
