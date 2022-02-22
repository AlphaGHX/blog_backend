package initialize

import (
	"blog/global"
	"blog/models"
)

func InitUser() {
	user := models.User{
		Username: global.CONFIG.User.Username,
		Password: global.CONFIG.User.Password,
	}
	global.GROM.Create(&user)
}
