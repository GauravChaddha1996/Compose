package commons

import "time"

const MAX_TIME = "2 Jan 3000 15:04:05"
const TimeFormat = "2 Jan 2006 15:04:05"

var MaxTime = time.Now()

func initTimeCoommons() {
	maxTime, err := time.ParseInLocation(TimeFormat, "2 Jan 3000 15:04:05", time.Now().Location())
	if InError(err) {
		panic(err)
	}
	MaxTime = maxTime
}

func ParseTime(value string) (time.Time, error) {
	return time.ParseInLocation(TimeFormat, value, time.Now().Location())
}
