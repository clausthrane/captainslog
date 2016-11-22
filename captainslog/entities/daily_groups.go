package entities

import (
	"time"
	"fmt"
)

func TodaysGroup() TaskGroupID {
	utc := time.Now().UTC()
	today := DateString(utc)
	return TaskGroupID(fmt.Sprintf("group:%s", today))
}

func GroupByDate(utc time.Time) TaskGroupID {
	today := DateString(utc)
	return TaskGroupID(fmt.Sprintf("group:%s", today))
}

func DateString(utc time.Time) string {
	return utc.Format("2006-01-02")
}

