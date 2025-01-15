package model

import "time"

type Message struct {
	ID             uint `gorm:"primaryKey;autoIncrement"`
	SenderID       uint `json:"sender_id"`
	Sender         *User
	RecipientID    *uint `json:"recipient_id"`
	Recipient      *User
	ConversationID uint	`json:"conversation_id"`
	Conversation   *Conversation
	Content        string
	Timestamp      time.Time `gorm:"autoCreateTime"`
	Status         string `gorm:"default:'sent';size:50"`
}

func (Message) TableName() string {
	return "messages"
}
