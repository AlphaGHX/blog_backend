package main

import (
	"blog/global"
	"blog/initialize"
	"fmt"
	// "blog/test"
)

func main() {
	global.VIPER = initialize.Viper()

	global.PrivateKey, global.PublicKey = initialize.InitKey()

	global.GROM = initialize.Gorm()

	initialize.InitUser()

	// test.InsertTestData()

	Router := initialize.Routers()
	fmt.Printf("Listen:\t%s\n", global.CONFIG.Local.ListeningAddr)
	Router.Run(global.CONFIG.Local.ListeningAddr)

	sqlDB, _ := global.GROM.DB()
	defer sqlDB.Close()
}
