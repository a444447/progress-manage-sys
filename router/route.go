package router

import (
	"github.com/gin-gonic/gin"
	v1 "progress-manage-system/api/v1"
	"progress-manage-system/middware"
	"progress-manage-system/utils"
)

func InitRouter() {
	r := gin.New()
	r.Use(middware.Logger())
	r.Use(gin.Recovery())

	auth := r.Group("api/v1")
	auth.Use(middware.JwtToken())
	{
		//User Controller
		auth.POST("user/add", v1.AddUser)
		auth.POST("user/search", v1.GetUsers)
		auth.DELETE("user/:id", v1.DelUser)
		auth.PUT("user/:id", v1.UpdateUser)
		auth.POST("user/map", v1.MapTest)

		//Thesis Controller
		auth.POST("thesis/add", thesisController.Create)
		auth.POST("thesis/search", thesisController.Find)
		auth.DELETE("thesis/:id", thesisController.Delete)
		auth.PUT("thesis/:id", thesisController.Update)

		//file Controller
		auth.PUT("upload", fileController.Upload)
		auth.POST("download", fileController.Download)
		auth.DELETE("file/:id", fileController.Delele)

		//task Controller
		auth.POST("task/add", taskController.CreateTask)
		auth.POST("task/findall", taskController.FindAllTasks)
		auth.DELETE("task/:id", taskController.DeleteTasks)
		auth.PUT("task/:id", taskController.UpdateTasks)
		auth.POST("task/findbymap", taskController.FindByMap)
	}

	router := r.Group("api/v1")
	{
		router.POST("login", v1.Login)
	}
	r.Run(utils.HttpPort)
}
