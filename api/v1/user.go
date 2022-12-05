package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"progress-manage-system/model"
	"progress-manage-system/utils/errmsg"
	"progress-manage-system/utils/validator" //自定义的validator包
	"strconv"
)

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User

	//绑定结构体
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	//数据校验是否通过
	if code := validator.Validate(data); code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		err := model.CreateUser(&data)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    errmsg.ErrDataTrans(data, err),
			"message": errmsg.ErrTrans(err),
		})
	}
}

// GetUserById
func GetUserById(c *gin.Context) {

}

// 查找条件查找用户
func GetUsers(c *gin.Context) {
	//条件查询不知道具体参数数量，需要动态
	data, err := model.DataMapByRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errmsg.ErrTrans(err),
		})
	}
	users, err := model.GetUsers(data)
	c.JSON(http.StatusOK, gin.H{
		"data":    users,
		"message": errmsg.ErrTrans(err),
	})
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
	id, _ := strconv.Atoi(c.Param("id"))
	user, _ := model.GetUserById(id)
	code := model.DelUser(id)
	if code == errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    user,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// 更新用户
func UpdateUser(c *gin.Context) {
	//目前缺少validate步骤
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := model.DataMapByRequest(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": errmsg.ErrTrans(err),
		})
	}
	err = model.UpdateUser(id, data)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": errmsg.ErrTrans(err),
	})
}
