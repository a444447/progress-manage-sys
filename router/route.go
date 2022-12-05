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
		router.POST("user/search", v1.GetUsers)
		router.DELETE("user/:id", v1.DelUser)
		router.PUT("user/:id", v1.UpdateUser)
		router.POST("user/map", v1.MapTest)
		router.POST("login", v1.Login)
	}
	r.Run(utils.HttpPort)
}
