package entities

import (
	"encoding/json"
	"bytes"
)

type CatagoryID string

func (c CatagoryID) IsBlank() bool {
	return c == ""
}

func BlankCategory() CatagoryID {
	return CatagoryID("")
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

func (g *TaskGroup) Get(groupID TaskGroupID, idx int) *Task {
	tasks := g.Tasks[groupID]
	return tasks[idx]
}

func (g *TaskGroup) Remove(groupID TaskGroupID, idx int) {
	tasks := g.Tasks[groupID]
	switch  {
	case idx == 0 && len(tasks) <= 1:
		g.Tasks[groupID] = nil
	case idx == 0:
		g.Tasks[groupID] = tasks[idx:]
	case idx < len(tasks):
		g.Tasks[groupID] = append(tasks[0:idx - 1], tasks[idx:]...)
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

func (l TaskGroupIDList) Len() int {
    return len(l)
}
func (l TaskGroupIDList) Swap(i, j int) {
    l[i], l[j] = l[j], l[i]
}
func (l TaskGroupIDList) Less(i, j int) bool {
    return string(l[i]) < string(l[j])
}