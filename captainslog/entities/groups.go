package entities

import (
	"encoding/json"
	"bytes"
)


type CatagoryID string

func (c CatagoryID) IsBlank() bool {
	return c == ""
}


// ID to identify groups of tasks
type TaskGroupID string

type TaskGroupIDList []TaskGroupID

// Group of tasks and metadata
type TaskGroup struct {
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

func (g *TaskGroup) Remove(groupID TaskGroupID, idx int) {
	tasks := g.Tasks[groupID]
	switch  {
	case idx > len(tasks):
		g.Tasks[groupID] = append(tasks[0:idx-1],tasks[idx:]...)
	case idx == 0:
		g.Tasks[groupID] = tasks[idx:]
	}
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
