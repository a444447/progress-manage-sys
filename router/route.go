package router

import (
	"github.com/gin-gonic/gin"
	v1 "progress-manage-system/api/v1"
	"progress-manage-system/utils"
)

func InitRouter() {
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
	}
	r.Run(utils.HttpPort)
}
