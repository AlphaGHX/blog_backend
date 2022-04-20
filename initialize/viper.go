package initialize

import (
	"blog/global"
	"fmt"

	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	fmt.Println("INIT\t Config")
	v := viper.New()
	v.SetConfigFile("./config.toml")
	v.SetConfigType("toml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("ERROR\t reading configuration file: %v", err))
	}
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		panic(fmt.Errorf("ERROR\t parsing configuration file: %v", err))
	}
	fmt.Println("DONE\t Config")
	return v
}
