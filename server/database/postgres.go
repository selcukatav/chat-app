package database

import (
	"log"

	"github.com/selcukatav/chat-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Postgres() {
	dsn := "host=localhost user=postgres password=123qwe dbname=chat_app port=5432 sslmode=disable TimeZone=Europe/Berlin"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	log.Println("Connected to database!")

	err = db.AutoMigrate(&models.User{}, &models.Conversation{}, &models.Contact{}, &models.ConversationParticipant{}, &models.Message{}, &models.Notification{}, &models.Error{})
	if err != nil {
		log.Fatal("error occured while migrating")
	}
	log.Println("Database migrated successfully!")
	DB = db
}
