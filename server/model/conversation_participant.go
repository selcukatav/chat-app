package model

import "time"

type ConversationParticipant struct {
	ID             uint `gorm:"primaryKey"`
	UserID         uint `json:"user_id"`
	User           *User
	ConversationID uint `json:"conversation_id"`
	Conversation   *Conversation
	Role           string    `gorm:"default:'member';size:50"`
	JoinedAt       time.Time `gorm:"autoCreateTime"`
}

func (ConversationParticipant) TableName() string {
	return "conversation_participants"
}
