package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/selcukatav/chat-app/model"
)

func (h *Handler) ChatRooms(c echo.Context) error {

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		fmt.Println("Websocket upgrader error: ", err)
		return err
	}
	defer conn.Close()

	clients[conn] = true
	defer delete(clients, conn)

	for {
		var msg model.Message

		if err := conn.ReadJSON(&msg); err != nil {
			fmt.Println("Write error:", err)
			break
		}

		if err := h.saveMessagesToPSQL(msg); err != nil {
			fmt.Println("Database error: ", err)
			continue
		}

		broadcast <- msg

	}

	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Welcome to the chat room!",
	})
}

func (h *Handler) saveMessagesToPSQL(msg model.Message) error {
	if err := h.DB.Create(&msg).Error; err != nil {
		return err
	}
	return nil
}
