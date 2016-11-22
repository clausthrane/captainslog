package entities

import (
	"encoding/json"
	"context"
	"time"
	"github.com/pborman/uuid"
	"github.com/fatih/color"
	"fmt"
	"github.com/clausthrane/captainslog/captainslog/utils"
)

var white = color.New(color.FgHiWhite).SprintFunc()
var green = color.New(color.FgHiGreen).SprintFunc()
var red = color.New(color.FgHiRed).SprintFunc()
var cyan = color.New(color.FgCyan).SprintFunc()


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
	return t.StringWthIdx(-1)
}

func (t *Task) StringWthIdx(idx int) string {

	done := green("√")
	if ! t.Done {
		done = red("∞")
	}

	return fmt.Sprintf("[%d] %s (%s) time: %s \n %s",
		idx,
		done,
		cyan(t.Category),
		utils.PrettyPrint(t.TimeUsed),
		white(t.Description))
}

type TaskList []*Task

func (list TaskList) SumTime() time.Duration {
	d := time.Duration(0)
	for _, task := range list {
		d = d + task.TimeUsed
	}
	return d
}