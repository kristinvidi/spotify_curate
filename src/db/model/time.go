package model

import "time"

type TimeFormat string

const (
	TimeFormatPostgresDate      TimeFormat = "2006-01-02"
	TimeFormatPostgresTime      TimeFormat = "15:04:05.999999999"
	TimeFormatPostgresTimestamp TimeFormat = "2006-01-02 15:04:05.999999999"
)

func FormatPostgresTime(time time.Time, format TimeFormat) string {
	return time.Format(string(format))
}
