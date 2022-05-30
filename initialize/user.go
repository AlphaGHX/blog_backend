package initialize

import (
	"blog/global"
	"blog/models"
)

func InitUser() {
	user := models.User{
		Username: global.CONFIG.User.Username,
		Password: global.CONFIG.User.Password,
		Nickname: global.CONFIG.User.Nickname,
	}
	if global.GORM.Find(&user).RowsAffected == 0 {
		global.GORM.Create(&user)
	}
}
