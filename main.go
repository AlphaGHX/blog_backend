package main

import (
	"blog/global"
	"blog/initialize"
)

func main() {
	// global.BLOG_LIST = initialize.ReadBlog()
	global.VIPER = initialize.Viper()

	global.GROM = initialize.Gorm()

	Router := initialize.Routers()
	Router.Run(":12900")

	// blog1 := models.Blog{
	// 	Name:  "test",
	// 	Title: "Markdown测试",
	// 	Text:  "Markdown的一些测试",
	// }
	// global.GROM.Create(&blog1)

	sqlDB, _ := global.GROM.DB()
	defer sqlDB.Close()
}
