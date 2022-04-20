package global

import (
	"blog/config"
	"blog/models"
	"crypto/rsa"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	BLOG_LIST  map[string]models.Blog
	GROM       *gorm.DB
	CONFIG     config.Config
	VIPER      *viper.Viper
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)
