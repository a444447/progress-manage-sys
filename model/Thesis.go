package model

import (
	"gorm.io/gorm"
)

type Thesis struct {
	gorm.Model
	Title     string `gorm:"type:varchar(32);not null" json:"title" validate:"required"`
	Desc      string `gorm:"type:varchar(400)" json:"desc,omitempty"`
	StuID     string `gorm:"type:varchar(32);not null" json:"stu_id" validate:"required"`
	Stu       User   `gorm:"foreignKey:StuID;references:IdentityID"`
	TeacherID string `gorm:"type:varchar(32);not null" json:"teacher_id" validate:"required"`
	Teacher   User   `gorm:"foreignKey:TeacherID;references:IdentityID"`
}

type ThesisService interface {
	//TODO 缺少分页查找
	Create(thesis *Thesis) error
	Delete(id int) (*Thesis, error)
	Update(id int, data map[string]interface{}) error
	FindByMap(data map[string]interface{}) ([]Thesis, error)
}

type ThesisRepository interface {
	CreateThesis(thesis *Thesis) error
	DeleteThesis(thesis *Thesis) error
	UpdateThesis(id int, data map[string]interface{}) error
	FindThesis(data map[string]interface{}) ([]Thesis, error)
	FindAll() ([]Thesis, error)
	FindById(id int) (*Thesis, error)
}

//
//// 录入毕业论文信息
//func (t *Thesis) loadThesis(data map[string]interface{}) error {
//
//	return nil
//}
//
//// 查询某id下的所有论文信息
//func (t *Thesis) searchById(id string) ([]*Thesis, error) {
//
//	return nil, nil
//}
//
//// 查询某个毕业论文信息(by title)
//func (t *Thesis) searchByTitle(title string) ([]*Thesis, error) {
//	return nil, nil
//}
//
//// 删除毕业论文信息
//func (t *Thesis) delThesis(id int) (*Thesis, error) {
//
//	return nil, nil
//}
//
//// 更新毕业论文信息
//func (t *Thesis) updateThesis(id int, data map[string]interface{}) (*Thesis, error) {
//	return nil, nil
//}
