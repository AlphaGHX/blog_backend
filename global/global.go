package global

import (
	"blog/config"
	"crypto/rsa"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GORM        *gorm.DB
	CONFIG      config.Config
	VIPER       *viper.Viper
	LOG         *logrus.Logger
	PRIVATE_KEY *rsa.PrivateKey
	PUBLIC_KEY  *rsa.PublicKey
)
