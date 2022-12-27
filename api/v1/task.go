package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"progress-manage-system/model"
	"progress-manage-system/utils/ecode"
	"progress-manage-system/utils/myLog"
	"strconv"
)

type TaskController interface {
	CreateTask(c *gin.Context)
	FindAllTasks(c *gin.Context)
	DeleteTasks(c *gin.Context)
	UpdateTasks(c *gin.Context)
	FindByMap(c *gin.Context)
}

type taskController struct {
	taskServ model.TaskService
}

func (t *taskController) FindByMap(c *gin.Context) {
	var (
		data  map[string]interface{}
		err   error
		tasks []model.Task
	)
	data, err = model.DataMapByRequest(c)
	if err != nil {
		myLog.Logger.Error(err)
		c.JSON(http.StatusOK, ecode.Response(ecode.ErrMapData))
		return
	}
	tasks, err = t.taskServ.FindByMap(data)
	if err != nil {
		myLog.Logger.Error(err)
		c.JSON(http.StatusOK, ecode.Response(err))
		return
	}
	c.JSON(http.StatusOK, ecode.Response(err, tasks))

}

func (t *taskController) UpdateTasks(c *gin.Context) {
	var (
		id   int
		err  error
		data map[string]interface{}
	)
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, ecode.ErrParamData)
	}
	data, err = model.DataMapByRequest(c)
	if err != nil {
		c.JSON(http.StatusOK, ecode.ErrMapData)
		return
	}
	err = t.taskServ.Update(id, data)
	c.JSON(http.StatusOK, ecode.Response(err))

}

func (t *taskController) DeleteTasks(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, ecode.Response(ecode.ErrParamData))
	}
	if err = t.taskServ.Delete(id); err != nil {
		c.JSON(http.StatusOK, ecode.Response(err))
		myLog.Logger.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, ecode.Response(err))

}

func (t *taskController) CreateTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusOK, ecode.Response(ecode.ErrBindModel))
		myLog.Logger.Error(err.Error())
		return
	}
	err := t.taskServ.Create(&task)
	c.JSON(http.StatusOK, ecode.Response(err))
}

func (t *taskController) FindAllTasks(c *gin.Context) {
	var tasks []model.Task
	var err error
	tasks, err = t.taskServ.FindAll()
	if err != nil {
		c.JSON(http.StatusOK, ecode.Response(err))
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, ecode.Response(err, tasks))
}

func NewTaskController(t model.TaskService) TaskController {
	return &taskController{
		taskServ: t,
	}
}
