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

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))
	//Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	//Auth
	e.GET(config.APILogin, handler.Login)
	e.POST(config.APIRegister, handler.Register)

	g := e.Group("/api")
	g.Use(middlewares.Authorize)

	g.GET(config.APIConversationRoom, handler.ConversationRoom)
	//Users
	g.PATCH(config.APIUpdateUser, handler.UpdateUser)
	g.GET(config.APIGetUser, handler.GetUser)
	g.GET(config.APIListUsers, handler.ListUsers)
	g.DELETE(config.APIDeleteUser, handler.DeleteUser)

	//Friends
	g.GET(config.APIGetUserFriends, handler.GetFriends)
	g.GET(config.APISearchFriends, handler.ListUsers)
	g.DELETE(config.APIDeleteFriend, handler.DeleteFriend)
	g.POST(config.APIAddFriend, handler.AddFriend)

	//Conversation
	g.POST(config.APICreateConversation, handler.CreateConversation)
	g.GET(config.APIListConversation, handler.ListConversations)
	g.POST(config.APIAddConversationParticipants, handler.AddConversationParticipant)
	g.DELETE(config.APIDeleteConversationParticipants, handler.DeleteConversationParticipant)
	g.GET(config.APIListConversationParticipants, handler.ListConversationsParticipants)
	g.GET(config.APIListUserConversations, handler.ListUserConversations)

	return e
}
