package utils

import (
	"strconv"
	"time"
)

func StringAsI64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 0)
}

func StringAsI64WithDefault(s string, def int64) int64 {
	if value, err := strconv.ParseInt(s, 10, 0); err == nil {
		return value
	}
	return def
}

func StringAsDuration(s string, scale time.Duration) time.Duration {
	value := StringAsI64WithDefault(s, 0)
	return time.Duration(value) * scale
}
