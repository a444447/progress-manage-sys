package model

import (
	"gorm.io/gorm"
	"time"
)

type Report struct {
	gorm.Model
	Publisher string    `gorm:"type:varchar(32);not null" json:"publisher"`
	PubModel  User      `gorm:"foreignKey:Publisher;references:IdentityID"`
	Receiver  string    `gorm:"type:varchar(32);not null" json:"receiver"`
	RecModel  User      `gorm:"foreignKey:Receiver;references:IdentityID"`
	Deadline  time.Time `gorm:"type:date; not null" json:"deadline"`
	//表示任务的完成情况
	IsCompleted int    `gorm:"type:int;not null" json:"is_completed" validate:"required,oneof=0 1 2"`
	Content     string `gorm:"type:text;not null" json:"content"`
	Type        int    `gorm:"type:int;not null" json:"type" validate:"required,oneof=0 1"`
	Attachment  int    `gorm:"type:int;not null" json:"attachment"`
	AttachFile  File   `gorm:"foreignKey:Attachment"`
}

type ReportService interface {
}

type ReportRepository interface {
}
