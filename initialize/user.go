package initialize

import (
	"blog/global"
	"blog/models"
	"errors"

	"gorm.io/gorm"
)

func InitUser() {
	user := models.User{
		Username: global.CONFIG.User.Username,
		Password: global.CONFIG.User.Password,
	}
	global.GROM.First(&user)
	if errors.Is(global.GROM.First(&user).Error, gorm.ErrRecordNotFound) {
		global.GROM.Create(&user)
	}
}
