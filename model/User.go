package model

import (
	"gorm.io/gorm"
	"log"
	"progress-manage-system/utils/errmsg"
)

type User struct {
	gorm.Model
	IdentityID string `gorm:"type:varchar(32);not null;unique" json:"identityID" validate:"required,min=5,max=30"`
	PassWord   string `gorm:"type:varchar(100);not null" json:"passWord" validate:"required,min=10,max=30"`
	Name       string `gorm:"type:varchar(32);not null" json:"name" validate:"required,min=1,max=30"`
	Gender     string `gorm:"type:uint;not null" json:"gender" validate:"required,oneof=1 2"`
	Role       string `gorm:"type:varchar(20);not null;DEFAULT:0001" json:"role" validate:"required"`
}

// 查询用户是否存在

// 新增用户
func CreateUser(data *User) int {
	data.PassWord, err = hashAndSalt(data.PassWord)
	if err != nil {
		log.Fatalln("hash and salt fail !")
		return errmsg.ERROR
	}
	if err := db.Create(&data).Error; err != nil {
		log.Fatalln("create user fail!")
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 批量新增用户

// 查询单个用户
func GetUser(id int) (User, int) {
	var user User
	if err := db.Where("ID = ?", id).First(&user).Error; err != nil {
		return user, errmsg.ERROR
	}

	return user, errmsg.SUCCESS
}

// 查询用户列表(可用于模糊查询)

// 删除用户

// 更新用户

// 登陆验证
