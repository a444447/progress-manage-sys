package v1

import (
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
	var code int
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if code = validator.Validate(data); code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		code = model.CreateUser(&data)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// 查找用户
func GetUser(c *gin.Context) {
	var id int
	id, _ = strconv.Atoi(c.Param("id"))
	data, code := model.GetUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
