package repository

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"progress-manage-system/model"
)

type fileRepository struct {
	DB *gorm.DB
}

func (f *fileRepository) DeleteFile(id int) error {
	var file model.File
	err := f.DB.Where("id = ?", id).Delete(&file).Error
	if err != nil {
		return errors.Wrapf(err, "Error DeleteFile():")
	}
	return nil
}

func (f *fileRepository) GetFile(id int) (*model.File, error) {
	var file model.File
	err := f.DB.Where("id = ?", id).First(&file).Error
	return &file, err

}

func (f *fileRepository) UploadFileInfo(file *model.File) error {
	//TODO implement me
	return f.DB.Create(&file).Error
}

// new a instance of fileRepository
func NewFileRepository(db *gorm.DB) model.FileRepository {
	return &fileRepository{
		DB: db,
	}
}
