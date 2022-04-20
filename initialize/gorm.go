package initialize

import (
	"blog/models"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	fmt.Println("INIT\t Database")
	if db, err := gorm.Open(sqlite.Open("blog.db")); err != nil {
		panic(fmt.Errorf("ERROR\t connecting to database: %v", err))
	} else {
		fmt.Println("DONE\t Database")
		fmt.Println("INIT\t AutoMigrate")
		setUpMysql(db)
		fmt.Println("DONE\t AutoMigrate")
		return db
	}
}

func setUpMysql(db *gorm.DB) {
	if err := db.AutoMigrate(models.Blog{}, models.User{}); err != nil {
		panic(fmt.Errorf("ERROR\t AutoMigrate: %v", err))
	}
}
