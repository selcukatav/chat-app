package model

import "time"

type Conversation struct {
	ID        uint      `gorm:"primaryKey"`
	Subject   string   `gorm:"size:255" json:"subject"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (Conversation) TableName() string {
	return "conversations"
}
