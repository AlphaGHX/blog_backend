package system

import (
	"blog/global"
	"blog/models"
	"blog/models/response"
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

func (s *UserService) GetAdminInfo() (*response.GetAdminInfo, error) {
	var myLinks []models.MyLink
	result := global.GROM.Find(&myLinks)
	if result.Error != nil {
		return nil, errors.New("find myLinks Error")
	}

	var adminInfo models.User
	result = global.GROM.Find(&adminInfo)
	if result.Error != nil {
		return nil, errors.New("find adminInfo Error")
	}

	data := response.GetAdminInfo{
		Nickname: adminInfo.Nickname,
		Describe: adminInfo.Describe,
		MyLinks:  myLinks,
	}

	return &data, nil
}

// func (s *UserService) SetAdminInfo(response.GetAdminInfo) error {

// }
