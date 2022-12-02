package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"progress-manage-system/utils"
	"time"
)

var db *gorm.DB
var err error

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassport,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	unableConnDB(err)
	db.AutoMigrate(&User{}, &Thesis{})

	sqlDB, err := db.DB()
	unableConnDB(err)
	//SetMaxIdleConns 设置空闲连接池中连接最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}
