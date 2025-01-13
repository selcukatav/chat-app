package model

import "time"

type Contact struct {
	ID            uint `gorm:"primaryKey"`
	UserID        uint
	User          *User 
	ContactUserID uint
	ContactUser   *User    
	IsAccepted    bool      `gorm:"default:false"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
}

func (Contact) TableName() string {
	return "contacts"
}
