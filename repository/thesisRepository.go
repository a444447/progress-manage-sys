package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"progress-manage-system/model"
)

type thesisRepository struct {
	DB *gorm.DB
}

func NewThesisRepository(db *gorm.DB) model.ThesisRepository {
	return &thesisRepository{
		DB: db,
	}
}

func (t *thesisRepository) CreateThesis(thesis *model.Thesis) error {

	return t.DB.Create(&thesis).Error
}

func (t *thesisRepository) DeleteThesis(thesis *model.Thesis) error {

	return t.DB.Delete(&thesis).Error
}

func (t *thesisRepository) UpdateThesis(id int, data map[string]interface{}) error {
	var thesis model.Thesis

	err := t.DB.Model(&thesis).Where("id = ?", id).Updates(data).Error
	return err
}

// 条件查找
func (t *thesisRepository) FindThesis(data map[string]interface{}) ([]model.Thesis, error) {
	var thesis []model.Thesis
	//clause.Associations能起作用，但是preload(user)不行
	err := t.DB.Preload(clause.Associations, func(DB *gorm.DB) *gorm.DB {
		return DB.Omit("passWord")
	}).Where(data).Find(&thesis).Error

	return thesis, err
}

func (t *thesisRepository) FindAll() ([]model.Thesis, error) {
	var thesis []model.Thesis
	err := t.DB.Find(&thesis).Error

	return thesis, err
}

func (t *thesisRepository) FindById(id int) (*model.Thesis, error) {
	var thesis *model.Thesis
	err := t.DB.Where("ID = ?", id).First(thesis).Error

	return thesis, err
}
