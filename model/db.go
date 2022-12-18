package model

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"progress-manage-system/utils"
	"time"
)

var Db *gorm.DB
var err error

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassport,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return Db, errors.Wrapf(err, "error->InitDB")
	}
	Db.AutoMigrate(&User{}, &Thesis{}, &File{})

	sqlDB, err := Db.DB()
	unableConnDB(err)
	//SetMaxIdleConns 设置空闲连接池中连接最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return Db, nil
}
