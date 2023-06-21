package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type DateTime struct{}

// Parse : custom parser of RFC3339 time format (e.g.: 2006-01-02T15:04:05Z)
func DatetimeParse(datetimeStr string) (*time.Time, error) {
	layoutParts := strings.Split(datetimeStr, "T")
	dateParts := strings.Split(layoutParts[0], "-")

	year, err := strconv.Atoi(dateParts[0])
	if err != nil {
		fmt.Println("Error during conversion")
		return nil, err
	}

	month, err := strconv.Atoi(dateParts[1])
	if err != nil {
		fmt.Println("Error during conversion")
		return nil, err
	}

	day, err := strconv.Atoi(dateParts[2])
	if err != nil {
		fmt.Println("Error during conversion")
		return nil, err
	}

	timeLayout := strings.Replace(layoutParts[1], "Z", "", 1)
	timeParts := strings.Split(timeLayout, ":")

	hours, err := strconv.Atoi(timeParts[0])
	if err != nil {
		fmt.Println("Error during conversion")
		return nil, err
	}

	minutes, err := strconv.Atoi(timeParts[1])
	if err != nil {
		fmt.Println("Error during conversion")
		return nil, err
	}

	seconds, err := strconv.Atoi(timeParts[2])
	if err != nil {
		fmt.Println("Error during conversion")
		return nil, err
	}

	parsedDatetime := time.Date(year, time.Month(month), day, hours, minutes, seconds, 0, time.UTC)

	return &parsedDatetime, nil
}
