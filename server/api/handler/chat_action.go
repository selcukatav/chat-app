package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/selcukatav/chat-app/database"
	"github.com/selcukatav/chat-app/model"
)

// websocket setup
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients sync.Map

//var broadcast = make(chan model.Message)

func (h *Handler) ChatAction(conn *websocket.Conn, conversationID string) error {
	defer conn.Close()
	channel := fmt.Sprintf("chat-room:%s", conversationID)

	clients.Store(conn, true)
	defer clients.Delete(conn)

	go h.SubscribeToRedis(channel)

	for {
		var msg model.Message

		if err := conn.ReadJSON(&msg); err != nil {
			fmt.Println("Read error: ", err)
			break
		}

		if err := h.saveMessagesToPSQL(msg); err != nil {
			fmt.Println("Error while saving to psql", err)
			continue
		}

		if err := h.publishMessageToRedis(channel, msg); err != nil {
			fmt.Println("Error while publishing msg to redis: ", err)
			continue
		}
	}
	return nil

}

func (h *Handler) saveMessagesToPSQL(msg model.Message) error {
	if err := h.DB.Create(&msg).Error; err != nil {
		return err
	}
	return nil
}

func (h *Handler) publishMessageToRedis(channel string, msg model.Message) error {
	client := database.Redis()

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	if err := client.Publish(context.Background(), channel, msgBytes).Err(); err != nil {
		return err
	}
	return nil
}

func (h *Handler) SubscribeToRedis(channel string) {
	client := database.Redis()

	pubsub := client.Subscribe(context.Background(), channel)

	go func() {
		for {
			msg, err := pubsub.ReceiveMessage(context.Background())
			if err != nil {
				fmt.Println("Subscribe err: ", err)
				break
			}
			var recievedMsg model.Message
			if err := json.Unmarshal([]byte(msg.Payload), &recievedMsg); err != nil {
				fmt.Println("Message unmarshall error: ", err)
				continue
			}
			clients.Range(func(key, value interface{}) bool {
				client := key.(*websocket.Conn)
				if err := client.WriteJSON(recievedMsg); err != nil {
					fmt.Println("write error: ", err)
					client.Close()
					clients.Delete(client)
				}
				return true
			})
		}
	}()
}
