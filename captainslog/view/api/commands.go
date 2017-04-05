package api

import (
	"context"
	"github.com/clausthrane/captainslog/captainslog/entities"
	"github.com/clausthrane/captainslog/captainslog/services"
	"time"
)

type Commands struct {
	ctx     context.Context
	service *services.TaskService
}

func NewCommands(ctx context.Context, service *services.TaskService) *Commands {
	return &Commands{ctx, service}
}

func (c *Commands) AddDoneTask(groupID entities.TaskGroupID, description string, category entities.CatagoryID, timeUsed time.Duration) (*entities.Task, error) {
	return c.addTask(groupID, description, category, timeUsed, true)
}

func (c *Commands) AddOpenTask(groupID entities.TaskGroupID, description string, category entities.CatagoryID) (*entities.Task, error) {
	return c.addTask(groupID, description, category, 0, false)
}

func (c *Commands) addTask(groupID entities.TaskGroupID, description string, category entities.CatagoryID, timeUsed time.Duration, done bool) (*entities.Task, error) {

	task := entities.NewTask(c.ctx, description, category, timeUsed, done)
	group, err := c.service.Load()
	if err != nil {
		return nil, err
	}
	group.Add(groupID, task)

	if e := c.service.Save(group); e != nil {
		return nil, e
	}

	return task, nil
}

func (c *Commands) RemoveTask(groupID entities.TaskGroupID, idx int) error {
	var err error
	if groups, err := c.service.Load(); err == nil {
		if err = groups.Remove(groupID, idx); err == nil {
			err = c.service.Save(groups)
		}
	}
	return err
}

