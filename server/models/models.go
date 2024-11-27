package models

import "time"

type User struct {
	UserID         uint                   `gorm:"primaryKey" json:"user_id"`
	Username       string                 `gorm:"unique" json:"username"`
	Email          string                 `gorm:"unique" json:"email"`
	UserPassword   string                 `gorm:"not null" json:"user_password"`
	ProfilePicture string                 `json:"profile_picture"`
	StatusMessage  string                 `json:"status_message"`
	CreatedAt      time.Time              `gorm:"autoCreateTime" json:"created_at"`

	// İlişkiler
	SentMessages      []Message           `gorm:"foreignKey:SenderID" json:"sent_messages"`
	ReceivedMessages  []Message           `gorm:"foreignKey:RecipientID" json:"received_messages"`
	Contacts          []Contact           `gorm:"foreignKey:UserID" json:"contacts"`
	Participations    []ConversationParticipant `gorm:"foreignKey:UserID" json:"participations"`
	Notifications     []Notification      `gorm:"foreignKey:UserID" json:"notifications"`
}


type Conversation struct {
	ConversationID uint                       `gorm:"primaryKey" json:"conversation_id"`
	Subject        string                     `json:"subject"`
	CreatedAt      time.Time                  `gorm:"autoCreateTime" json:"created_at"`

	// İlişkiler
	Participants     []ConversationParticipant `gorm:"foreignKey:ConversationID" json:"participants"`
	Messages         []Message                 `gorm:"foreignKey:ConversationID" json:"messages"`
}


type ConversationParticipant struct {
	ParticipantID  uint        `gorm:"primaryKey" json:"participant_id"`
	ConversationID uint        `gorm:"not null" json:"conversation_id"`
	UserID         uint        `gorm:"not null" json:"user_id"`
	Role           string      `gorm:"default:member" json:"role"`
	JoinedAt       time.Time   `gorm:"autoCreateTime" json:"joined_at"`

	// İlişkiler
	Conversation    Conversation `gorm:"foreignKey:ConversationID;references:ConversationID" json:"conversation"`
	User            User         `gorm:"foreignKey:UserID;references:UserID" json:"user"`
}


type Message struct {
	MessageID      uint        `gorm:"primaryKey" json:"message_id"`
	SenderID       uint        `gorm:"not null" json:"sender_id"`
	RecipientID    *uint       `json:"recipient_id"` // Nullable (Pointer kullanımı)
	Content        string      `gorm:"not null" json:"content"`
	Timestamp      time.Time   `gorm:"autoCreateTime" json:"timestamp"`
	ConversationID uint        `gorm:"not null" json:"conversation_id"`
	Status         string      `gorm:"default:sent" json:"status"`

	// İlişkiler
	Sender          User         `gorm:"foreignKey:SenderID;references:UserID" json:"sender"`
	Recipient       *User        `gorm:"foreignKey:RecipientID;references:UserID" json:"recipient"`
	Conversation    Conversation `gorm:"foreignKey:ConversationID;references:ConversationID" json:"conversation"`
}


type Contact struct {
	ContactID     uint      `gorm:"primaryKey" json:"contact_id"`
	UserID        uint      `gorm:"not null" json:"user_id"`
	ContactUserID uint      `gorm:"not null" json:"contact_user_id"`
	IsAccepted    bool      `gorm:"default:false" json:"is_accepted"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`

	// İlişkiler
	User        User `gorm:"foreignKey:UserID;references:UserID" json:"user"`
	ContactUser User `gorm:"foreignKey:ContactUserID;references:UserID" json:"contact_user"`
}


type Notification struct {
	NotificationID   uint      `gorm:"primaryKey" json:"notification_id"`
	UserID           uint      `gorm:"not null" json:"user_id"`
	Content          string    `gorm:"not null" json:"content"`
	Timestamp        time.Time `gorm:"autoCreateTime" json:"timestamp"`
	IsRead           bool      `gorm:"default:false" json:"is_read"`
	NotificationType string    `gorm:"not null" json:"notification_type"`

	User User `gorm:"foreignKey:UserID;references:UserID" json:"user"`
}


type Error struct {
	LogID             uint   `gorm:"primaryKey" json:"log_id"`
	ResponseCode      int16  `json:"rc"`
	Message           string `json:"message"`
	Detail            string `json:"detail"`
	ExternalReference string `json:"ext_ref"`
}
