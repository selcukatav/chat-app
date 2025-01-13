package model

import "time"

type Notification struct {
	ID               uint `gorm:"primaryKey"`
	UserID           uint
	User             *User
	Content          string
	Timestamp        time.Time `gorm:"autoCreateTime"`
	IsRead           bool      `gorm:"default:false"`
	NotificationType string    `gorm:"size:50;"`
}

func (Notification) TableName() string {
	return "notifications"
}
