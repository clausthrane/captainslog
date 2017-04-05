package api

import (
	"context"
	"github.com/clausthrane/captainslog/captainslog/services"
	"github.com/clausthrane/captainslog/captainslog/entities"
	"os"
	"log"
)

var logger = log.New(os.Stdout, " ", log.Ldate | log.Ltime | log.Lshortfile)

type Queries struct {
	ctx     context.Context
	service *services.TaskService
}

func NewQueries(ctx context.Context, service *services.TaskService) *Queries {
	return &Queries{ctx, service}
}

func (q *Queries) ListTasks(groupId entities.TaskGroupID, category entities.CatagoryID) (entities.TaskList, error) {
	group, err := q.service.Load()
	if err != nil {
		return nil, err
	}

	tasks := group.Tasks[groupId]
	if ! category.IsBlank() {
		res := tasks[:0]
		for _, task := range tasks {
			if task.Category == category {
				res = append(res, task)
			}
		}
		tasks = res
	}

	return tasks, nil
}

func (q *Queries) ListGroups() (entities.TaskGroupIDList, error) {
	if group, err := q.service.Load(); err != nil {
		return nil, err
	} else {
		keys := make([]entities.TaskGroupID, 0, len(group.Tasks))
		for key := range group.Tasks {
			keys = append(keys, key)
		}
		//sort.Sort(keys)
		return keys, nil
	}
}

func (c *Queries) GetTask(groupID entities.TaskGroupID, idx int) (*entities.Task, error) {
	var err error
	if groups, err := c.service.Load(); err == nil {
		return groups.Get(groupID, idx), nil
	}
	return nil, err
}