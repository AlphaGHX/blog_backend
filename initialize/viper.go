package initialize

import (
	"blog/global"
	"fmt"

	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	fmt.Println("INFO\t读取配置...")
	v := viper.New()
	v.SetConfigFile("./config.toml")
	v.SetConfigType("toml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("ERROR\t配置读取时错误: %v", err))
	}
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		panic(fmt.Errorf("ERROR\t配置写入内存时错误: %v", err))
	}
	fmt.Println("DONE\t读取配置成功")
	return v
}
