package commons

import "time"

const MAX_TIME = "2 Jan 3000 15:04:05"
const TimeFormat = "2 Jan 2006 15:04:05"

func ParseTime(value string) (time.Time, error) {
	return time.ParseInLocation(TimeFormat, value, time.Now().Location())
}

func MaxTime() (time.Time, error) {
	return time.ParseInLocation(TimeFormat, "2 Jan 3000 15:04:05", time.Now().Location())
}
