package services

import (
	"time"
	"fmt"
)

var (
	LONG_DATE_TIME = "2006-01-02 15:04:05"
	TIME_ZONE = "15:04:05"
)

func DateTimeNow() string {
	return time.Now().Format(LONG_DATE_TIME)
}

func GetTimeZoneString(dateTime string) string {

	return ConvertStringToDateTime(dateTime).Format(TIME_ZONE)
}

func ConvertStringToDateTime (dateTime string) time.Time {
	backToTime, err := time.Parse(LONG_DATE_TIME, dateTime)
	if err != nil {
		fmt.Println("error parsing time")
	}

	return backToTime
}