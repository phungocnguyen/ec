package services

import (
	"time"
	"fmt"
)

var (
	LONG_DATE_TIME = "Mon Jan 2 15:04:05 MST 2006"
	TIME_ZONE = "15:04:05 MST"
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