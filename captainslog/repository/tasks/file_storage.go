package task_repository

import (
	"io/ioutil"
	"github.com/clausthrane/captainslog/captainslog/entities"
	"github.com/clausthrane/captainslog/captainslog/config"
)

var logger = config.Logger

type filebacked_tasks struct {
	ProjectRoot string
}

func NewFileBackedTaskRepository(projectRoot string) TaskRepository {
	return &filebacked_tasks{projectRoot}
}

func (f *filebacked_tasks) Save(projectID entities.ProjectID, data []byte) error {
	filename := f.ProjectFile(projectID)
	logger.Debug("Saving to", filename)
	return ioutil.WriteFile(filename, data, 0644)
}

func (f *filebacked_tasks) Load(projectID entities.ProjectID) ([]byte, error) {
	filename := f.ProjectFile(projectID)
	return ioutil.ReadFile(filename)
}

func (f *filebacked_tasks) ProjectFile(projectID entities.ProjectID) string {
	return f.ProjectRoot + "/" + string(projectID) + ".json"

}
