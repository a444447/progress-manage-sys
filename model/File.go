package model

import (
	"gorm.io/gorm"
	"mime/multipart"
)

//以id唯一表示每个文件，无论是图像，文档
//存储文件的路径尽量是选择根路径，以便如果以后更换云存储方，不必更改整个数据库的文件存储路径字段

type File struct {
	gorm.Model
	Name     string `gorm:"type:varchar(250);not null" json:"name"`
	SavePath string `gorm:"type:varchar(250);not null" json:"save_path"`
	//Type表示上传附近者的身份，目前暂时用1,2,3分别表示教师、学生、管理员，未来可以增加一个关系对应表
	Type uint8 `gorm:"type:int;not null" json:"type"`
}

type FileService interface {
	Upload(fd *multipart.FileHeader) (string, error)
	Download(id int) error
	Delete(id int) error
}

type FileRepository interface {
	UploadFileInfo(file *File) error
	GetFile(id int) (*File, error)
	DeleteFile(id int) error
}
