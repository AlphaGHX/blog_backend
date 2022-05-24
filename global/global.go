package global

import (
	"blog/config"
	"crypto/rsa"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GROM        *gorm.DB
	CONFIG      config.Config
	VIPER       *viper.Viper
	FILE_LOG    *logrus.Logger
	STD_LOG     *logrus.Logger
	PRIVATE_KEY *rsa.PrivateKey
	PUBLICK_KEY *rsa.PublicKey
)
