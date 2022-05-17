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
	result := global.GROM.Find(&data, map[string]interface{}{"username": user.Username, "password": user.Password})
	if result.RowsAffected == 0 {
		return errors.New("VerifyUser Error")
	}
	return result.Error
}

func (s *UserService) GetAdminInfo() ([]models.MyLink, error) {
	var data []models.MyLink
	result := global.GROM.Find(&data)
	if result.Error != nil {
		return nil, errors.New("find MyLinks ERROR")
	}
	return data, nil
}
