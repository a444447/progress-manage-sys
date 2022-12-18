package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"progress-manage-system/model"
	"progress-manage-system/utils/ecode"
	"strconv"
)

type FileController interface {
	Upload(c *gin.Context)
	Download(c *gin.Context)
	Delele(c *gin.Context)
}

type fileController struct {
	fileServ model.FileService
}

func (f *fileController) Delele(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, ecode.Response(ecode.ErrParamData))
		return
	}
	err = f.fileServ.Delete(id)
	c.JSON(http.StatusOK, ecode.Response(err))
}

func (f *fileController) Download(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(http.StatusOK, ecode.Response(ecode.ErrMapData))
		return
	}
	err = f.fileServ.Download(id)
	c.JSON(http.StatusOK, ecode.Response(err))
}

func (f *fileController) Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	fmt.Println(file.Filename)
	path, err := f.fileServ.Upload(file)
	if err != nil {
		c.JSON(http.StatusOK, ecode.Response(err))
		return
	}
	c.JSON(http.StatusOK, ecode.Response(err, path))
}

func NewFileService(fs model.FileService) FileController {
	return &fileController{
		fileServ: fs,
	}
}
