package main

import (
	"blog/global"
	"blog/initialize"
	// "blog/temp"
)

func main() {
	global.VIPER = initialize.Viper()

	global.GROM = initialize.Gorm()

	initialize.InitUser()

	// temp.InsertTestData()

	Router := initialize.Routers()
	Router.Run("localhost:12900")

	sqlDB, _ := global.GROM.DB()
	defer sqlDB.Close()
}
