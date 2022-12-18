package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
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

// 遍历获得gin请求的所有参数(post方法)
func DataMapByRequest(c *gin.Context) (map[string]interface{}, error) {
	//我们先解析Form(这里解析的是post方法)
	contentType := c.GetHeader("Content-Type")
	data := make(map[string]interface{})
	switch contentType {
	case "application/x-www-form-urlencoded":
		fmt.Println("form type")
		if err := c.Request.ParseForm(); err != nil {
			return nil, errors.Wrapf(err, "error->DataMapByRequest #%d", 1)
		}
		//c.PostForm("") //为了调用initFormCache
		for k, v := range c.Request.PostForm {
			if len(v) > 1 {
				errMsg := fmt.Sprintf("[%+v]设置了%d次, 但只能设置一次", k, len(v))
				return nil, errors.New(errMsg)
			}
			if k == "ID" || k == "id" {
				idInt, _ := strconv.Atoi(c.PostForm(k))
				data[k] = idInt
			} else {
				data[k] = c.PostForm(k)
			}
		}
		break
	case "application/json":
		fmt.Println("json type")
		if err := c.BindJSON(&data); err != nil {
			return nil, errors.Wrapf(err, "error->DataMapByRequest #%d", 2)
		}
		break
	}

	return data, nil
	//get请求通过url传递参数的解析方式
	//for k, _ := range c.Request.URL.Query() {
	//	fmt.Printf("key:%+v value:%+v\n", k, c.Query(k))
	//}
}
