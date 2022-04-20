package system

import (
	"blog/global"
	"blog/models"
	"errors"
)

type UserService struct {
}

func (s *UserService) VerifyUser(user models.User) error {
	var data models.User
	result := global.GROM.Where(map[string]interface{}{"username": user.Username, "password": user.Password}).Find(&data)
	if result.RowsAffected == 0 {
		return errors.New("VerifyUser Error")
	}
	return result.Error
}
