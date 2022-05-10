package main

import (
	"blog/global"
	"blog/initialize"
	"fmt"
)

func main() {
	// 加载配置文件
	global.VIPER = initialize.Viper()

	// 加载密钥
	global.PrivateKey, global.PublicKey = initialize.InitKey()

	// 初始化ORM
	global.GROM = initialize.Gorm()

	// 初始化管理员
	initialize.InitUser()

	// 创建本地数据目录
	initialize.InitFile()

	// 初始化路由
	Router := initialize.Routers()
	fmt.Printf("Listen:\t%s\n", global.CONFIG.Local.ListeningAddr)
	Router.Run(global.CONFIG.Local.ListeningAddr)

	sqlDB, _ := global.GROM.DB()
	defer sqlDB.Close()
}
