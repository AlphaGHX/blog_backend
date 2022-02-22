package system

import (
	"blog/global"
	"blog/models"
	"errors"

	"gorm.io/gorm"
)

type UserService struct {
}

func (s *UserService) VerifyUser(user models.User) error {
	var data models.User
	result := global.GROM.Where(map[string]interface{}{"username": user.Username, "password": user.Password}).First(&data)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("VerifyUser Error")
	}
	return result.Error
}
