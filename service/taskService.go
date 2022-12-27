package service

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"progress-manage-system/model"
	"progress-manage-system/utils/ecode"
)

type taskService struct {
	taskRepo model.TaskRepository
}

func (t taskService) FindByMap(data map[string]interface{}) ([]model.Task, error) {

	return t.taskRepo.FindByMap(data)

}

func (t taskService) Create(task *model.Task) error {

	return t.taskRepo.CreateTask(task)
}

func (t taskService) Delete(id int) error {
	var task model.Task
	var err error
	task, err = t.FindById(id)
	if err != nil {
		return errors.Wrapf(err, "taskServ error->Delete")
	}
	return t.taskRepo.DeleteTask(&task)
}

func (t taskService) FindById(id int) (model.Task, error) {
	task, err := t.taskRepo.FindById(id)
	return *task, err
}

func (t taskService) Update(id int, data map[string]interface{}) error {
	//id是否存在
	if _, err := t.taskRepo.FindById(id); err != nil {
		return ecode.ErrTaskNotFound
	}

	//如果更新的json中带有了profile
	var ProfileList model.TaskProfileList
	if profiles, ok := data["profile"]; ok {
		for _, profile := range profiles.([]interface{}) {
			var taskProfile model.TaskProfile
			p := profile.(map[string]interface{}) //断言为map类型
			marshal, _ := json.Marshal(p)
			fmt.Println(marshal)
			json.Unmarshal(marshal, &taskProfile)
			ProfileList = append(ProfileList, taskProfile)
		}
		data["profile"] = ProfileList
	}
	return t.taskRepo.UpdateTask(id, data)
}

func (t taskService) FindAll() ([]model.Task, error) {
	tasks, err := t.taskRepo.FindAll()
	return tasks, err
}

func NewTaskService(t model.TaskRepository) model.TaskService {
	return &taskService{
		taskRepo: t,
	}
}
