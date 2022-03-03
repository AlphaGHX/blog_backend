package global

import (
	"blog/config"
	"blog/models"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	BLOG_LIST map[string]models.Blog
	GROM      *gorm.DB
	CONFIG    config.Config
	VIPER     *viper.Viper
	KEY_FILE	[]byte
)
