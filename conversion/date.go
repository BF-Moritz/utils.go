package conversion

import (
	"fmt"
	"time"
)

type TimeType time.Time

func TimeToDateString(t time.Time) string {
	return t.Format("2006-01-02")
}

func TimeToTimeString(t time.Time) string {
	return t.Format("15:04:05")
}

func TimeToDateTimeString(t time.Time) string {
	return fmt.Sprintf("%s %s", TimeToDateString(t), TimeToTimeString(t))
}
