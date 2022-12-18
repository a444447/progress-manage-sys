package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"progress-manage-system/model"
	"progress-manage-system/utils/ecode"
	"strconv"
)

type ThesisController interface {
	Create(c *gin.Context)
	Find(c *gin.Context) //条件查询
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

type thesisController struct {
	thesisService model.ThesisService
}

func (t *thesisController) Create(c *gin.Context) {
	var thesis model.Thesis
	if err := c.ShouldBindJSON(&thesis); err != nil {
		resp := ecode.Response(err)
		c.JSON(http.StatusOK, resp)
		return
	}
	err := t.thesisService.Create(&thesis)
	resp := ecode.Response(err)
	c.JSON(http.StatusOK, resp)
}

// Find 此函数为条件查找的接口
func (t *thesisController) Find(c *gin.Context) {
	var resp *ecode.ResponseJSON
	data, err := model.DataMapByRequest(c)
	if err != nil {
		resp = ecode.Response(ecode.ErrMapData)
		c.JSON(http.StatusOK, resp)
		return
	}
	//service返回的是Ok,不是nil
	thesis, err := t.thesisService.FindByMap(data)
	if !ecode.Cause(err).Equal(ecode.Ok) {
		c.JSON(http.StatusOK, ecode.Response(err))
		return
	}
	c.JSON(http.StatusOK, ecode.Response(err, thesis))

}

func (t *thesisController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, ecode.Response(err))
		return
	}
	thesis, err := t.thesisService.Delete(id)
	if !ecode.Cause(err).Equal(ecode.Ok) {
		c.JSON(http.StatusOK, ecode.Response(err))
		return
	}
	c.JSON(http.StatusOK, ecode.Response(err, thesis))
}

func (t *thesisController) Update(c *gin.Context) {

	//TODO 需要解决id不存在的问题
	var (
		id   int
		err  error
		data map[string]interface{}
	)
	id, err = strconv.Atoi(c.Param("id"))
	data, err = model.DataMapByRequest(c)
	if err != nil {
		c.JSON(http.StatusOK, ecode.Response(err))
		return
	}
	err = t.thesisService.Update(id, data)
	c.JSON(http.StatusOK, ecode.Response(err))

}

// New an instance of thesis controller
func NewThesisController(thesisServ model.ThesisService) ThesisController {
	return &thesisController{
		thesisService: thesisServ,
	}
}
