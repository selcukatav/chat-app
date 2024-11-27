package middlewares

import (
	"errors"

	"github.com/selcukatav/chat-app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Authentication(postgres *gorm.DB, username, password string) error {

	var user models.User
	result := postgres.Where("username=?", username).First(&user)
	if errors.Is(result.Error,gorm.ErrRecordNotFound){
		return errors.New("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.UserPassword),[]byte(password))
	if err!=nil{
		return errors.New("wrong password")
	}

	return nil
}
