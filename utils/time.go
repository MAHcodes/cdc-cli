package utils

import "time"

var timeFormat = "15:04 - 2006-01-02"

func FormatUnixTime(timestamp int) string {
	return time.Unix(int64(timestamp), 0).Format(timeFormat)
}
