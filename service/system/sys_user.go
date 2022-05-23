package system

import (
	"blog/global"
	"blog/models"
	"blog/models/request"
	"blog/models/response"
	"errors"
)

type UserService struct {
}

func (s *UserService) VerifyUser(user models.User) error {
	var data models.User
	result := global.GROM.Find(&data, map[string]interface{}{"username": user.Username, "password": user.Password})
	if result.RowsAffected == 0 {
		return errors.New("VerifyUser error")
	}
	return result.Error
}

func (s *UserService) GetAdminInfo() (*response.GetAdminInfo, error) {
	var myLinks []models.MyLink
	result := global.GROM.Find(&myLinks)
	if result.Error != nil {
		return nil, errors.New("find myLinks error")
	}

	var adminInfo models.User
	result = global.GROM.Find(&adminInfo)
	if result.Error != nil {
		return nil, errors.New("find adminInfo error")
	}

	data := response.GetAdminInfo{
		Nickname: adminInfo.Nickname,
		Describe: adminInfo.Describe,
		MyLinks:  myLinks,
	}

	return &data, nil
}

func (s *UserService) SetAdminInfo(adminInfo request.SetAdminInfo) error {
	var data models.User
	global.GROM.Find(&data)
	{
		data.Describe = adminInfo.Describe
		data.Nickname = adminInfo.Nickname
	}
	result := global.GROM.Save(&data)
	if result.Error != nil {
		return errors.New("update user error")
	}

	// 删除表中所有数据
	global.GROM.Exec("DELETE FROM my_links")

	result = global.GROM.Create(adminInfo.MyLinks)
	if result.Error != nil {
		return errors.New("create my_links error")
	}

	return nil
}
