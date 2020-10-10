package dateUtils

import "time"

const (
	DateLayout = "2006-01-02T15:04:05Z"
	DBDateLayout = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(DateLayout)
}

func GetNowDbFormat() string {
	return GetNow().Format(DBDateLayout)
}
