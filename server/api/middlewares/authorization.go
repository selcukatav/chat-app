package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
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
		"exp":      time.Now().Add(time.Hour * 24),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return accessToken, nil

}

func SetCookie()
