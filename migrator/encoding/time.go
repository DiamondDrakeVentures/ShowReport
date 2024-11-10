package encoding

import "strings"

const (
	FormatTime        = "3:04 PM"
	FormatDate        = "1/2/2006"
	FormatDateTime    = "1/2/2006 3:04 PM"
	FormatDateTimeSec = "1/2/2006 3:04:05 PM"
	FormatTimestamp   = "1/2/2006 3:04:05 PM"
)

func timeFormat(id string) string {
	switch strings.ToLower(id) {
	case "date":
		return FormatDate

	case "datetime", "date_time":
		return FormatDateTime

	case "datetimesec", "datetime_sec", "date_time_sec":
		return FormatDateTimeSec

	case "timestamp":
		return FormatTimestamp

	case "time":
		fallthrough

	default:
		return FormatTime
	}
}
