package model

import "gorm.io/gorm"

type Thesis struct {
	gorm.Model
	Title     string
	Desc      string
	StuID     string
	Stu       User `gorm:"foreignKey:StuID;references:IdentityID"`
	TeacherID string
	Teacher   User `gorm:"foreignKey:TeacherID;references:IdentityID"`
}

// 录入毕业论文信息

// 查询某指导老师管理下的所有论文信息

// 查询某个毕业论文信息(by title)

// 删除毕业论文信息

// 更新毕业论文信息
