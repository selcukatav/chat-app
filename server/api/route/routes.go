package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/selcukatav/chat-app/api/handler"
	"github.com/selcukatav/chat-app/api/middlewares"
	"github.com/selcukatav/chat-app/config"
	"github.com/selcukatav/chat-app/database"
	_ "github.com/selcukatav/chat-app/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           Discord Clone API
// @version         1.0
// @description     Discord-like API Documentation.
// @license.name   Apache 2.0
// @license.url    http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:3000
// @BasePath  /
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

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Welcome to Discord Clone API"})
	})
	//Auth
	e.GET(config.APILogin, handler.Login)
	e.POST(config.APIRegister, handler.Register)
	//users
	e.PUT(config.APIUpdateUser, handler.UpdateUser)
	e.GET(config.APIGetUser, handler.GetUser)
	e.GET(config.APIListUsers, handler.ListUsers)
	e.DELETE(config.APIDeleteUser, handler.DeleteUser)
	//friends
	e.GET(config.APIGetUserFriends, handler.GetFriends)
	e.GET(config.APISearchFriends, handler.ListUsers)
	e.DELETE(config.APIDeleteFriend, handler.DeleteUser)
	e.POST(config.APIAddFriend,handler.AddFriend)

	ChatRooms(g)

	return e
}

func ChatRooms(g *echo.Group) {
	g.Use(middlewares.Authorize)
	g.GET("/rooms", handler.Rooms)
}
