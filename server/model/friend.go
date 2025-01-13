package model

import "time"

type Friend struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	UserID    uint `json:"user_id"`
	User      *User
	FriendID  uint `json:"friend_id"`
	Friend    *User
	Status    string    `gorm:"default:'pending'"`
	Timestamp time.Time `gorm:"autoCreateTime"`
}

func (Friend) TableName() string {
	return "friends"
}
