package database

import (
	"log"

	"github.com/selcukatav/chat-app/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Postgres() *gorm.DB {
	dsn := "host=localhost user=postgres password=123qwe dbname=chat_app port=5432 sslmode=disable TimeZone=Europe/Berlin"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	log.Println("Connected to database!")
	db = db.Debug()
	// db.Migrator().DropTable(&model.User{},
	// 	&model.Conversation{},
	// 	&model.Contact{},
	// 	&model.ConversationParticipant{},
	// 	&model.Message{},
	// 	&model.Notification{},
	// 	)

	err = db.AutoMigrate(
		&model.User{},
		&model.Conversation{},
		&model.Contact{},
		&model.ConversationParticipant{},
		&model.Message{},
		&model.Notification{},
		&model.Friend{},
		&model.Error{})
	if err != nil {
		log.Fatal("error occured while migrating")
	}

	log.Println("Database migrated successfully!")
	DB = db
	return db

}
