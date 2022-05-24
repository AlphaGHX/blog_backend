package initialize

import (
	"blog/global"
	"blog/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	if db, err := gorm.Open(sqlite.Open("blog.db")); err != nil {
		global.STD_LOG.Errorln("connecting to database: ", err)
		panic(err)
	} else {
		setUpSQLite(db)
		return db
	}
}

func setUpSQLite(db *gorm.DB) {
	if err := db.AutoMigrate(models.Blog{}, models.User{}, models.MyLink{}); err != nil {
		global.STD_LOG.Errorln("connecting to database: ", err)
		panic(err)
	}
}
