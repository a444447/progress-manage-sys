package model

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"progress-manage-system/utils/baseTime"
)

type Task struct {
	gorm.Model
	Publisher  string          `gorm:"type:varchar(32);not null" json:"publisher" validate:"required"`
	PubModel   User            `gorm:"foreignKey:Publisher;references:IdentityID"`
	Receiver   string          `gorm:"type:varchar(32);not null" json:"receiver" validate:"required"`
	RecModel   User            `gorm:"foreignKey:Receiver;references:IdentityID"`
	Profile    TaskProfileList `gorm:"json" json:"profile" validate:"required"`
	Type       int             `gorm:"type:int;not null" json:"type" validate:"required,oneof=0 1"` //任务的类型，0为计划表，1为周期汇报任务
	Attachment *int            `gorm:"type:int" json:"attachment,omitempty"`                        //设置为指针，这样零值是nil而不会是0
	AttachFile File            `gorm:"foreignKey:Attachment" json:"attach_file,omitempty"`
}

type TaskProfile struct {
	Title       string            `json:"title" validate:"required"`
	Content     string            `json:"content"`
	StartTime   baseTime.BaseTime `json:"start_time" validate:"required"`
	EndTime     baseTime.BaseTime `json:"end_time" validate:"required"`
	IsCompleted int               `json:"is_completed" validate:"required,oneof=0 1"` //0为未完成，1为已完成
	IsOvertime  int               `json:"is_overtime" validate:"required,oneof=0 1"`  //0为未逾期，1未已逾期
}

type TaskProfileList []TaskProfile

func (t TaskProfileList) Value() (driver.Value, error) {
	return json.Marshal(t)
}

// 负责解析进结构体，写入到数据库
func (t *TaskProfileList) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &t)
}

type TaskService interface {
	Create(task *Task) error
	Delete(id int) error
	FindById(id int) (Task, error)
	Update(id int, data map[string]interface{}) error
	FindAll() ([]Task, error)
	FindByMap(data map[string]interface{}) ([]Task, error) //条件查找
	//查找还有未完成事项的任务
	//查询未超时的任务
	//查询超时的
}

type TaskRepository interface {
	CreateTask(task *Task) error
	DeleteTask(task *Task) error
	FindById(id int) (*Task, error)
	UpdateTask(id int, data map[string]interface{}) error
	FindAll() ([]Task, error)
	FindByMap(data map[string]interface{}) ([]Task, error)     //条件查找
	FindJSONQuery(data map[string]interface{}) ([]Task, error) //如果条件查找需要对profile内的内容查找，调用此Find
}
