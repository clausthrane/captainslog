package entities

import (
	"encoding/json"
	"context"
	"time"
	"bytes"
)

// ID type for tasks
type TaskID string

// ID to identify groups of tasks
type TaskGroupID string

// Basic state for a task
type Task struct {
	Id          TaskID `json:"task_id"`
	User        UserID `json:"user"`
	Description string `json:"description"`
	Created     time.Time `json:"created"`
}

func NewTask(ctx context.Context, description string) *Task {
	return &Task{
		"foobar",
		ctx.Value("user").(UserID),
		description,
		time.Now(),
	}
}

func (t *Task) marshal() []byte {
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}

	return b
}


// Group of tasks
type TaskGroup struct {
	//Name string `json:"name"`
	//Tasks []Task `json:"tasks"`
	Tasks map[TaskGroupID][]*Task `json:"tasks"`
}

func NewTaskGroup() *TaskGroup {
	return &TaskGroup{
		make(map[TaskGroupID][]*Task),
	}
}

func (g *TaskGroup) Add(groupID TaskGroupID, task *Task) {
	group := g.Tasks[groupID]
	if group == nil {
		group = []*Task{task}
	} else {
		group = append(group, task)
	}
	g.Tasks[groupID] = group
}

func (taskList *TaskGroup) Marshal() ([]byte, error) {
	return json.Marshal(taskList)
}

func NewTaskGroupFromData(data []byte) (*TaskGroup, error) {
	var group TaskGroup
	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&group)
	if err != nil {
		return nil, err
	}
	return &group, nil
}
