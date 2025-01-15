package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/selcukatav/chat-app/model"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan model.Message)

func (h *Handler) HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Println("Write Error:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
