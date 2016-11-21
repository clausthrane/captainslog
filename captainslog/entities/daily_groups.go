package entities

import (
	"time"
	"fmt"
)


func TodaysGroup() TaskGroupID {
	utc := time.Now().UTC()
	today := utc.Format("2006-01-02")
	return TaskGroupID(fmt.Sprintf("group:%s", today))
}

func GroupByDate(utc time.Time) TaskGroupID {
	today := utc.Format("2006-01-02")
	return TaskGroupID(fmt.Sprintf("group:%s", today))
}

