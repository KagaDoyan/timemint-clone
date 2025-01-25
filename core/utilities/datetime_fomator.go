package utilities

import (
	"fmt"
	"time"
)

var dateTimeFormat = "2006-01-02 15:04:05"

func NewDateTimeFormatToString(datetime time.Time) string {
	return datetime.Format(dateTimeFormat)
}

func DateTimeFormat(datetime string) (time.Time, bool) {
	// Define the layout of your input string
	layout := "2006-01-02 15:04:05 PM"
	t, err := time.Parse(layout, datetime)
	if err != nil {
		fmt.Println("Error:", err)
		return time.Now(), false
	}
	return t, true
}
