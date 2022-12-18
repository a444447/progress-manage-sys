package service

import (
	"progress-manage-system/model"
	"progress-manage-system/utils/ecode"
)

type thesisService struct {
	thesisRepository model.ThesisRepository
}

func (t *thesisService) Create(thesis *model.Thesis) error {
	//TODO 判断是否重复Create
	//TODO 判断外键关系是否满足
	//调用User的CheckUser方法，注意该方法返回nil表示用户不存在
	e1 := model.CheckUser(thesis.StuID)
	e2 := model.CheckUser(thesis.TeacherID)
	if e1 == nil || e2 == nil {
		if e1 == nil {
			return ecode.ErrThesisStuId
		}
		return ecode.ErrThesisTeacherId
	}
	err := t.thesisRepository.CreateThesis(thesis)
	if err != nil {
		return ecode.ErrCreate
	}

	return ecode.Ok
}

func (t *thesisService) Delete(id int) (*model.Thesis, error) {
	var thesis *model.Thesis
	//获得对应id的实体
	tt, err := t.thesisRepository.FindById(id)
	if err != nil {
		//如果findbyid失败，直接返回错误
		return tt, ecode.Cause(err)
	}
	err = t.thesisRepository.DeleteThesis(tt)
	if err != nil {
		//如果delete失败，直接返回错误
		return thesis, ecode.ErrDelete
	}

	return thesis, ecode.Ok

}

func (t *thesisService) Update(id int, data map[string]interface{}) error {
	if err := t.thesisRepository.UpdateThesis(id, data); err != nil {
		return ecode.ErrUpdate
	}

	return ecode.Ok

}

func (t *thesisService) FindByMap(data map[string]interface{}) ([]model.Thesis, error) {
	thesis, err := t.thesisRepository.FindThesis(data)
	if err != nil {
		return thesis, ecode.Cause(err)
	}
	return thesis, ecode.Ok
}

// New a thesisService instance
func NewThesisService(thesisRepo model.ThesisRepository) model.ThesisService {
	return &thesisService{
		thesisRepository: thesisRepo,
	}
}
