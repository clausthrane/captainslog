package utils

import (
	"time"
	"strings"
	"fmt"
	"math"
)

var day = 24 * time.Hour

var weekdays = map[string]int{
	"mon": 0,
	"tue": 1,
	"wed": 2,
	"thu": 3,
	"fri": 4,
	"sat": 5,
	"sun": 6,
}

func ParseDate(dateString string) time.Time {
	t, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		panic(err)
	}
	return t
}

func WeekdayToDate(w string) time.Time {
	dayidx := weekdays[strings.ToLower(w)]
	utc := time.Now().UTC()
	return time.Date(utc.Year(), utc.Month(), dayidx, 0, 0, 0, 0, time.UTC)
}


func PrettyPrint(duration time.Duration) string {
	hours := duration / time.Hour
	minutes := (duration - (hours * time.Hour)) / time.Minute
	return fmt.Sprintf("%d:%d", hours, minutes)
}

func DatesInThisWeek() []time.Time {
	now := time.Now().UTC()
	thisWeekday := time.Weekday(math.Max(0, float64(now.Weekday())-1))
	duration := -day * time.Duration(thisWeekday)
	start := now.Add(duration)
	dates := make([]time.Time, 7)
	for i :=0 ; i<7; i++{
		dates[i] = start.Add(time.Duration(i) * day)
	}
	return dates
}

func DateString(t time.Time) string {
	return t.Format("2006-01-02")
}