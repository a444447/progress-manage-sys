package model

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// 数据库连接失败
func unableConnDB(err error) {
	if err != nil {
		log.Println("unable to connect database!")
		//os.Exit(1)
	}
}

// hashAndSalt 加密Hash And Salt
func hashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// 验证密码
func ComparePassword(hashedPwd, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return false
	}
	return true
}
