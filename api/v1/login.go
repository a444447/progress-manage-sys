package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"progress-manage-system/model"
	"progress-manage-system/utils/errmsg"
)

// 此login函数主要是for admin
// 可以修改datamap函数，使得我们可以用json格式，也可以用form形式
func Login(c *gin.Context) {
	var user model.User
	var code int
	var token string
	params := []string{"identityID", "passWord"}
	data, _ := model.DataMapByRequest(c)
	//检查参数是否匹配
	for _, p := range params {
		if _, ok := data[p]; !ok {
			c.JSON(http.StatusOK, gin.H{
				"status":  errmsg.ErrorLoginParams,
				"message": errmsg.GetErrMsg(errmsg.ErrorLoginParams),
			})
			return
		}
	}
	user.IdentityID, _ = data["identityID"].(string)
	user.PassWord, _ = data["passWord"].(string)
	//数据校验

	//登陆检查
	user, code = model.CheckLogin(user.IdentityID, user.PassWord)
	if code == errmsg.SUCCESS {
		token = "test token"
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
