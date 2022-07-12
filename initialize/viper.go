package initialize

import (
	"blog/global"

	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("./config.toml")
	v.SetConfigType("toml")
	if err := v.ReadInConfig(); err != nil {
		global.LOG.Errorln("reading configuration file: ", err)
		panic(err)
	}
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		global.LOG.Errorln("parsing configuration file: ", err)
		panic(err)
	}
	return v
}
