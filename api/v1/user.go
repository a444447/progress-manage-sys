package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"progress-manage-system/model"
	. "progress-manage-system/utils/ecode"
	"progress-manage-system/utils/validator" //自定义的validator包
	"strconv"
)

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var resp *ResponseJSON
	//绑定结构体
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	//数据校验是否通过
	if err := validator.Validate(data); err != nil {
		resp = Response(err)
		c.JSON(http.StatusOK, resp)
	} else {
		err := model.CreateUser(&data)
		resp = Response(err, data)
		c.JSON(http.StatusOK, resp)
	}
}

// GetUserById
func GetUserById(c *gin.Context) {

}

// 条件查找用户
func GetUsers(c *gin.Context) {
	var err error
	var data map[string]interface{}
	var users []model.User
	var resp *ResponseJSON
	//条件查询不知道具体参数数量，需要动态
	data, err = model.DataMapByRequest(c)
	if err != nil {
		fmt.Printf("err: %+v", err)
	}
	users, err = model.GetUsers(data)
	resp = Response(err, users)
	c.JSON(http.StatusOK, resp)
}

// map test
func MapTest(c *gin.Context) {
	data, err := model.DataMapByRequest(c)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// 删除用户
func DelUser(c *gin.Context) {
	var err error
	var user model.User
	var resp *ResponseJSON
	id, _ := strconv.Atoi(c.Param("id"))
	user, err = model.GetUserById(id)
	if err != nil {
		//id查询不到就不用执行下面的删除操作了
		resp = Response(err)
		c.JSON(http.StatusOK, resp)
		return
	}
	err = model.DelUser(id)
	if err != nil {
		fmt.Printf("error:%+v", err)
		return
	}
	resp = Response(err, user)
	c.JSON(http.StatusOK, resp)
}

// 更新用户
func UpdateUser(c *gin.Context) {
	var resp *ResponseJSON
	//目前缺少validate步骤
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := model.DataMapByRequest(c)
	if err != nil {
		fmt.Printf("error:%+v", err)
		return
	}
	err = model.UpdateUser(id, data)
	if err != nil {
		fmt.Printf("error:%+v", err)
		resp = Response(err)
		c.JSON(http.StatusOK, resp)
		return
	}
	resp = Response(err)
	c.JSON(http.StatusOK, resp)

}
