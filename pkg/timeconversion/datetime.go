package timeconversion

import "time"

func DateTimeToTimestamp(dataTime string) int64 {
	timestamp, _ := time.ParseInLocation("2006-01-02 15:04:05", dataTime, time.Local)
	timeUnix := timestamp.Unix()
	return timeUnix
}
