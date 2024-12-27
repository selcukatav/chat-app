package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/selcukatav/chat-app/api/handler"
	"github.com/selcukatav/chat-app/api/middlewares"
	"github.com/selcukatav/chat-app/database"
)

func New() *echo.Echo {
	e := echo.New()
	db := database.Postgres()
	handler := &handler.Handler{
		DB: db,
	}
	g := e.Group("/rooms")

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	e.POST("/login", handler.Login)
	e.POST("/register", handler.Register)
	ChatRooms(g)

	return e
}

func ChatRooms(g *echo.Group) {
	g.Use(middlewares.Authorize)
	g.GET("/rooms", handler.Rooms)
}
