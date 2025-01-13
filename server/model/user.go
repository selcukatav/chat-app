package model

import "time"

type User struct {
	ID             uint      `gorm:"primaryKey" json:"user_id"`
	Username       string    `gorm:"unique;size:100" json:"username"`
	Email          string    `gorm:"unique;size:255" json:"email"`
	UserPassword   string    `gorm:"size:255" json:"user_password"`
	ProfilePicture string    `gorm:"type:text" json:"profile_picture"`
	StatusMessage  string    `gorm:"size:255" json:"status_message"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}

func (User) TableName() string {
	return "users"
}
