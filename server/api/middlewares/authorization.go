package middlewares

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

//var claims jwt.MapClaims

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
