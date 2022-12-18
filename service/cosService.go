package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
	"progress-manage-system/utils"
)

/*
此service用于上传附件至腾讯云cos的对象存储中
*/

var COS_URL_FORMAT string = "https://%s-%s.cos.%s.myqcloud.com"

type cosService struct {
	c *cos.Client
}

// 初始化cosService实例
func NewCosService() cosService {
	u, _ := url.Parse(fmt.Sprintf(COS_URL_FORMAT, utils.COS_BUCKET_NAME, utils.COS_APP_ID, utils.COS_REGION))
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  utils.COS_SECRET_ID,
			SecretKey: utils.COS_SECRET_KEY,
		},
	})
	return cosService{
		c: c,
	}
}

// 上传对象
func (cs *cosService) UploadObj(file *multipart.FileHeader, savePath string) error {
	//TODO 使用加密方式把文件名字变成独一的,防止重复
	fd, err := file.Open()
	if err != nil {
		return errors.Wrapf(err, "COS upload fail: file can not open")
	}
	defer fd.Close()
	_, err = cs.c.Object.Put(context.Background(), savePath, fd, nil)
	if err != nil {
		return errors.Wrapf(err, "COS upload fail: file can not upload to bucket")
	}
	//TODO Put可以返回uri
	return nil
}

// 下载对象
func (cs *cosService) DownloadObj(target, destination string) error {
	opt := &cos.MultiDownloadOptions{
		ThreadPoolSize: 5,
	}
	_, err := cs.c.Object.Download(context.Background(), target, destination, opt)
	if err != nil {
		return errors.Wrapf(err, "COS download fail")
	}
	return nil
}

// 删除对象
func (cs *cosService) DeleteObj(target string) error {
	_, err := cs.c.Object.Delete(context.Background(), target)
	if err != nil {
		return errors.Wrapf(err, "COS delete fail")
	}
	return nil
}
