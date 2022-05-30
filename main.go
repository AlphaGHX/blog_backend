package main

import (
	"blog/global"
	"blog/initialize"
)

func main() {
	// 初始化日志系统
	global.FILE_LOG, global.STD_LOG = initialize.Logrus()

	// 加载配置文件
	global.VIPER = initialize.Viper()

	// 加载密钥
	global.PRIVATE_KEY, global.PUBLIC_KEY = initialize.InitKey()

	// 初始化ORM
	global.GORM = initialize.Gorm()

	// 初始化管理员
	initialize.InitUser()

	// 创建本地数据目录
	initialize.InitFile()

	// 初始化路由
	Router := initialize.Routers()
	global.STD_LOG.Infoln("Listen: ", global.CONFIG.Local.ListeningAddr)
	global.FILE_LOG.Infoln("Listen: ", global.CONFIG.Local.ListeningAddr)
	Router.Run(global.CONFIG.Local.ListeningAddr)

	sqlDB, _ := global.GORM.DB()
	defer sqlDB.Close()
}
