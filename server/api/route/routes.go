package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/selcukatav/chat-app/api/handler"
)

func New() *echo.Echo {
	e := echo.New()
	h := &handler.Handler{}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	e.POST("/login", h.Login)
	e.POST("/register", h.Register)

	return e
}
