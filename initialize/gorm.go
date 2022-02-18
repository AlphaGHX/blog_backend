package initialize

import (
	"blog/global"
	"blog/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	fmt.Println("INFO\t连接数据库...")
	dbConfig := global.CONFIG.Mysql
	dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ")/" + dbConfig.Dbname + "?charset=" + dbConfig.Charset + "&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		panic(fmt.Errorf("ERROR\t数据库连接时错误: %v", err))
	} else {
		fmt.Println("DONE\t连接数据库成功")
		fmt.Println("INFO\t初始化表...")
		setUpMysql(db)
		fmt.Println("DONE\t初始化表成功")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		return db
	}
}

func setUpMysql(db *gorm.DB) {
	if err := db.AutoMigrate(models.Blog{}); err != nil {
		panic(fmt.Errorf("ERROR\t数据库初始化表时错误: %v", err))
	}
}
