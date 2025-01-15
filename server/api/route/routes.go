package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/selcukatav/chat-app/api/handler"
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
	//g := e.Group("/rooms")

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))
	//Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	//websocket conc start
	go handler.HandleMessages()

	e.GET(config.APIChatRooms, handler.ChatRooms)

	//Auth
	e.GET(config.APILogin, handler.Login)
	e.POST(config.APIRegister, handler.Register)

	//Users
	e.PATCH(config.APIUpdateUser, handler.UpdateUser)
	e.GET(config.APIGetUser, handler.GetUser)
	e.GET(config.APIListUsers, handler.ListUsers)
	e.DELETE(config.APIDeleteUser, handler.DeleteUser)

	//Friends
	e.GET(config.APIGetUserFriends, handler.GetFriends)
	e.GET(config.APISearchFriends, handler.ListUsers)
	e.DELETE(config.APIDeleteFriend, handler.DeleteFriend)
	e.POST(config.APIAddFriend, handler.AddFriend)

	//Conversation
	e.POST(config.APICreateConversation, handler.CreateConversation)
	e.GET(config.APIListConversation, handler.ListConversations)
	e.POST(config.APIAddConversationParticipants, handler.AddConversationParticipant)
	e.DELETE(config.APIDeleteConversationParticipants, handler.DeleteConversationParticipant)
	e.GET(config.APIListConversationParticipants, handler.ListConversationsParticipants)

	return e
}
