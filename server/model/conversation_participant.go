package model

import "time"

type ConversationParticipant struct {
	ID             uint `gorm:"primaryKey"`
	UserID         uint
	User           *User
	ConversationID uint
	Conversation   *Conversation
	Role           string    `gorm:"default:'member';size:50"`
	JoinedAt       time.Time `gorm:"autoCreateTime"`
}

func (ConversationParticipant) TableName() string {
	return "conversation_participants"
}
