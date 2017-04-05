package entities

import "fmt"

type WeeklyReport map[CatagoryID][]*Task

func NewWeeklyReport() *WeeklyReport {
	rep := make(WeeklyReport)
	return &rep
}

func (r *WeeklyReport) Index(category CatagoryID, task *Task) {
	items := r[task.Category]
	if items == nil {
		items = make([]*Task, 0)
	}
	items = append(items, task)
	r[task.Category] = items
}

func (r WeeklyReport) String() string {
	s := ""
	for key, value := range r {
		s = s + key.String()
		for _, task := range value {
			s = s + fmt.Sprintf("- %s \n", task.Description)
		}
	}
	return s
}
