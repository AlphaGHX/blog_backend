package initialize

import (
	"blog/models"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	fmt.Println("INFO\t连接数据库...")
	if db, err := gorm.Open(sqlite.Open("blog.db")); err != nil {
		panic(fmt.Errorf("ERROR\t数据库连接时错误: %v", err))
	} else {
		fmt.Println("DONE\t连接数据库成功")
		fmt.Println("INFO\t初始化表...")
		setUpMysql(db)
		fmt.Println("DONE\t初始化表成功")
		return db
	}
}

func setUpMysql(db *gorm.DB) {
	if err := db.AutoMigrate(models.Blog{}, models.User{}); err != nil {
		panic(fmt.Errorf("ERROR\t数据库初始化表时错误: %v", err))
	}
}
