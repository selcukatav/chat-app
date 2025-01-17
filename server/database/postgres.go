package database

import (
	"fmt"
	"log"
	"os"

	"github.com/selcukatav/chat-app/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Postgres() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Berlin",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	fmt.Println("DSN:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	log.Println("Connected to database!")
	db = db.Debug()
	// db.Migrator().DropTable(&model.User{},
	// 	&model.Conversation{},
	//
	// 	&model.ConversationParticipant{},
	// 	&model.Message{},
	// 	&model.Notification{},
	// 	)

	err = db.AutoMigrate(
		&model.User{},
		&model.Conversation{},
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
