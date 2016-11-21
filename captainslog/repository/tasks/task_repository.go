package task_repository

import "github.com/clausthrane/captainslog/captainslog/entities"

type TaskRepository interface {
	/**
	Registre a new task
	 */
	Save(entities.ProjectID, []byte) error

	/*
	List all Tasks for the given group
	 */
	Load(entities.ProjectID) ([]byte, error)
}
