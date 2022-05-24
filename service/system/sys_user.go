package system

import (
	"blog/global"
	"blog/models"
	"blog/models/request"
	"blog/models/response"
)

type UserService struct {
}

func (s *UserService) VerifyUser(user models.User) error {
	var data models.User
	result := global.GROM.Find(&data, map[string]interface{}{"username": user.Username, "password": user.Password})
	if result.RowsAffected == 0 {
		global.STD_LOG.Errorln("UserService.VerifyUser RowsAffected == 0")
		global.FILE_LOG.Errorln("UserService.VerifyUser RowsAffected == 0")
	}
	return result.Error
}

func (s *UserService) GetAdminInfo() (*response.GetAdminInfo, error) {
	var myLinks []models.MyLink
	result := global.GROM.Find(&myLinks)
	if result.Error != nil {
		global.STD_LOG.Errorln("UserService.GetAdminInfo Find myLinks", result.Error.Error())
		global.FILE_LOG.Errorln("UserService.GetAdminInfo Find myLinks", result.Error.Error())
		return nil, result.Error
	}

	var adminInfo models.User
	result = global.GROM.Find(&adminInfo)
	if result.Error != nil {
		global.STD_LOG.Errorln("UserService.GetAdminInfo Find adminInfo", result.Error.Error())
		global.FILE_LOG.Errorln("UserService.GetAdminInfo Find adminInfo", result.Error.Error())
		return nil, result.Error
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
		global.STD_LOG.Errorln("UserService.SetAdminInfo Save", result.Error.Error())
		global.FILE_LOG.Errorln("UserService.SetAdminInfo Save", result.Error.Error())
		return result.Error
	}

	result = global.GROM.Exec("DELETE FROM my_links")
	if result.Error != nil {
		global.STD_LOG.Errorln("UserService.SetAdminInfo Exec", result.Error.Error())
		global.FILE_LOG.Errorln("UserService.SetAdminInfo Exec", result.Error.Error())
		return result.Error
	}

	result = global.GROM.Create(adminInfo.MyLinks)
	if result.Error != nil {
		global.STD_LOG.Errorln("UserService.SetAdminInfo Exec", result.Error.Error())
		global.FILE_LOG.Errorln("UserService.SetAdminInfo Exec", result.Error.Error())
		return result.Error
	}

	return result.Error
}
