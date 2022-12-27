package repository

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"progress-manage-system/model"
)

type taskRepository struct {
	DB *gorm.DB
}

func (t taskRepository) FindJSONQuery(data map[string]interface{}) ([]model.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t taskRepository) FindByMap(data map[string]interface{}) ([]model.Task, error) {
	var tasks []model.Task
	err := t.DB.Preload("PubModel").Where(data).Find(&tasks).Error
	if err != nil {
		return nil, errors.Wrapf(err, "taskRepo error->FindByMap")
	}
	return tasks, err
}

func (t taskRepository) CreateTask(task *model.Task) error {
	return t.DB.Create(&task).Error
}

func (t taskRepository) DeleteTask(task *model.Task) error {
	return t.DB.Delete(&task).Error
}

func (t taskRepository) FindById(id int) (*model.Task, error) {
	var task model.Task
	err := t.DB.Preload(clause.Associations).Where("id = ?", id).First(&task).Error
	if err != nil {
		return &task, errors.Wrapf(err, "taskRepo error->FindById")
	}
	return &task, err

}

func (t taskRepository) UpdateTask(id int, data map[string]interface{}) error {
	//更新任务的时候，我们把更新profile与更新其他内容分开讨论，因为我们存储profile是以json数组的形式存在Mysql中的
	err := t.DB.Model(&model.Task{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return errors.Wrapf(err, "taskRepo err->UpdateTask")
	}

	return err
}

func (t taskRepository) FindAll() ([]model.Task, error) {
	var task []model.Task
	err := t.DB.Preload(clause.Associations).Find(&task).Error
	if err != nil {
		return task, errors.Wrapf(err, "taskRepo error->FindAll")
	}
	return task, err
}

// New an instance of taskRepo
func NewTaskRepository(db *gorm.DB) model.TaskRepository {
	return &taskRepository{
		DB: db,
	}
}
