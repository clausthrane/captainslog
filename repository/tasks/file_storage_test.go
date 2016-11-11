package task_repository

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/clausthrane/captainslog/entities"
)



func TestRepositoryCanStoreTasks(t *testing.T) {
	assert := assert.New(t)
	repo := NewFileBackedTaskRepository("/tmp")

	projectID := entities.ProjectID("Someproject")

	input_value := "hello"
	repo.Save(projectID, []byte(input_value))
	data, _ := repo.Load(projectID)

	assert.Equal(input_value, string(data))
}
