package model

import (
	"fmt"
	"gorm.io/gorm"
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
func CheckUser(iden string) int {
	var user User
	db.Select("id").Where("identity_id = ?", iden).First(&user)
	if user.ID > 0 {
		return errmsg.ErrorUserExisted
	}
	return errmsg.SUCCESS
}

// 新增用户
func CreateUser(data *User) error {

	//先检测是否以及存在用户
	if code := CheckUser(data.IdentityID); code == errmsg.ErrorUserExisted {
		return fmt.Errorf("record has existed")
	}
	data.PassWord, err = hashAndSalt(data.PassWord)
	if err != nil {
		return fmt.Errorf("hash and salt fail ")
	}
	if err := db.Create(&data).Error; err != nil {
		return fmt.Errorf("create user fail")
	}
	return nil
}

// 批量新增用户

// 查询单个用户
func GetUserById(id int) (User, error) {
	var user User
	if err := db.Where("ID = ?", id).First(&user).Error; err != nil {
		return user, fmt.Errorf("get user by id fail:%s", err.Error())
	}

	return user, nil
}

// 查询用户列表(条件查询查询)
func GetUsers(data map[string]interface{}) ([]User, error) {
	var users []User
	for k, v := range data {
		fmt.Printf("%+v\n", k)
		fmt.Printf("%+v\n", v)
	}
	if err := db.Where(data).Find(&users).Error; err != nil {
		return nil, fmt.Errorf("get users fail:%s", err.Error())
	}
	return users, nil

}

// 删除用户
func DelUser(id int) int {
	var user User
	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		fmt.Println(err.Error())
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 更新用户
func UpdateUser(id int, data map[string]interface{}) error {
	var user User
	if err := db.Model(&user).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

// 登陆验证
func CheckLogin(username, password string) (User, int) {
	var user User
	db.Where("identity_id = ?", username).First(&user)
	//账号不存在
	if user.ID == 0 {
		return user, errmsg.ErrorUserNotExist
	}
	//密码错误
	if !ComparePassword(user.PassWord, password) {
		return user, errmsg.ErrorPasswordWrong
	}
	//无法以此权限登陆

	return user, errmsg.SUCCESS
}
