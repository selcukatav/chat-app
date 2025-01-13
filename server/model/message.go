package model

import "time"

type Message struct {
	ID             uint `gorm:"primaryKey;autoIncrement"`
	SenderID       uint
	Sender         *User
	RecipientID    *uint
	Recipient      *User
	ConversationID uint
	Conversation   *Conversation
	Content        string
	Timestamp      time.Time `gorm:"autoCreateTime"`
	Status         string `gorm:"default:'sent';size:50"`
}

func (Message) TableName() string {
	return "messages"
}
