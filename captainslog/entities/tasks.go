package entities

import (
	"encoding/json"
	"context"
	"time"
	"github.com/pborman/uuid"
	"fmt"
)

// ID type for tasks
type TaskID uuid.UUID

// Basic state for a task
type Task struct {
	Id          TaskID         `json:"task_id"`
	User        UserID         `json:"user"`
	Description string         `json:"description"`
	Category    CatagoryID     `json:"catagory"`
	TimeUsed    time.Duration  `json:"minutes"`
	Done        bool           `json:"Done"`
	Created     time.Time      `json:"created"`
}

func NewTask(ctx context.Context, description string, category CatagoryID, timeUsed time.Duration, done bool) *Task {
	if category == "" {
		category = CatagoryID("N/A")
	}
	return &Task{
		TaskID(uuid.NewRandom()),
		ctx.Value("user").(UserID),
		description,
		category,
		timeUsed,
		done,
		time.Now().UTC(),
	}
}

func (t *Task) marshal() []byte {
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}

	return b
}

func (t *Task) String() string {
	done := "√"
	if ! t.Done {
		done = "∞"
	}
	return fmt.Sprintf("%s (%s) Minutes: %d \n %s \n",
		done,
		t.Category,
		int(t.TimeUsed.Minutes()),
		t.Description)
}

type TaskList []*Task

func (list TaskList) SumTime() time.Duration {
	d := time.Duration(0)
	for _, task := range list {
		d = d + task.TimeUsed
	}
	return d
}