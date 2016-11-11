package services

import (
	"github.com/clausthrane/captainslog/entities"
	"github.com/clausthrane/captainslog/repository/tasks"
)

func NewTaskService(dao task_repository.TaskRepository, projectID entities.ProjectID) *TaskService {
	return &TaskService{dao, projectID}
}

type TaskService struct {
	dao       task_repository.TaskRepository
	projectID entities.ProjectID
}

func (s *TaskService) Save(taskGroup *entities.TaskGroup) error {
	if data, err := taskGroup.Marshal(); err != nil {
		return err
	} else {
		return s.dao.Save(s.projectID, data)

	}
}

func (s *TaskService) Load() (*entities.TaskGroup, error) {
	if data, err := s.dao.Load(s.projectID); err != nil {
		return nil, err
	} else {
		if len(data) == 0 {
			return entities.NewTaskGroup(), nil
		} else {
			return entities.NewTaskGroupFromData(data)
		}
	}
}