package main

import (
	"blog/global"
	"blog/initialize"
	// "blog/temp"
)

func main() {
	global.VIPER = initialize.Viper()

	global.GROM = initialize.Gorm()

	Router := initialize.Routers()
	Router.Run(":12900")

	// temp.InsertTestData()

	sqlDB, _ := global.GROM.DB()
	defer sqlDB.Close()
}
