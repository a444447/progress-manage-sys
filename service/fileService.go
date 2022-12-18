package service

import (
	"fmt"
	"github.com/pkg/errors"
	"mime/multipart"
	"progress-manage-system/model"
	"strconv"
	"time"
)

type fileService struct {
	fileRepo model.FileRepository
	cosServ  cosService
}

func (f *fileService) Upload(fd *multipart.FileHeader) (string, error) {
	//为了避免名字重复，使用filename + time.Now().Unix的形式进行拼接

	path := fmt.Sprintf("%s-%s", strconv.Itoa(int(time.Now().Unix())), fd.Filename)
	if err := f.cosServ.UploadObj(fd, path); err != nil {
		return path, err
	}
	file := model.File{
		Name:     fd.Filename,
		SavePath: path,
		Type:     1,
	}
	if err := f.fileRepo.UploadFileInfo(&file); err != nil {
		return path, err
	}
	return path, nil
}

func (f *fileService) Download(id int) error {

	file, err := f.fileRepo.GetFile(id)
	if err != nil {
		return err
	}
	savePath := file.SavePath
	dst := file.Name
	if err := f.cosServ.DownloadObj(savePath, dst); err != nil {
		return err
	}
	return nil

}

func (f *fileService) Delete(id int) error {
	file, err := f.fileRepo.GetFile(id)
	if err != nil {
		return err
	}
	savaPath := file.SavePath
	if err = f.cosServ.DeleteObj(savaPath); err != nil {
		return errors.Wrapf(err, "Error Delete():")
	}
	if err = f.fileRepo.DeleteFile(id); err != nil {
		return errors.Wrapf(err, "Error Delete():")
	}
	return nil
}

func NewFileService(cs cosService, f model.FileRepository) model.FileService {
	return &fileService{
		cosServ:  cs,
		fileRepo: f,
	}
}
