package dates

import (
	"errors"
	"fmt"
	"time"
)

var (
	dateFormat = "2006-01-02"
	isoFormat  = "%04d-%02d-%02d"
)

func GetAbbrWeekdayName(weekday time.Weekday) (string, error) {
	var abbrWeekday string
	switch weekday {
	case 0:
		abbrWeekday = "Sun"
	case 1:
		abbrWeekday = "Mon"
	case 2:
		abbrWeekday = "Tue"
	case 3:
		abbrWeekday = "Wed"
	case 4:
		abbrWeekday = "Thu"
	case 5:
		abbrWeekday = "Fri"
	case 6:
		abbrWeekday = "Sat"
	default:
		return "", errors.New("No valid weekday given")
	}
	return abbrWeekday, nil
}

func GetAbbrWeekdayNameGerman(weekday time.Weekday) (string, error) {
	var abbrWeekday string
	switch weekday {
	case 0:
		abbrWeekday = "So"
	case 1:
		abbrWeekday = "Mo"
	case 2:
		abbrWeekday = "Di"
	case 3:
		abbrWeekday = "Mi"
	case 4:
		abbrWeekday = "Do"
	case 5:
		abbrWeekday = "Fr"
	case 6:
		abbrWeekday = "Sa"
	default:
		return "", errors.New("No valid weekday given")
	}
	return abbrWeekday, nil
}

// GetDateFirstOfMonth returns a time.Time object
// For a given year and month (as ints) it calculates the
// first date in the month.
func GetDateFirstOfMonth(year, month int) (time.Time, error) {
	return time.Parse(dateFormat, fmt.Sprintf("%04d-%02d-%02d", year, month, 1))
}

// GetDateLastOfMonth returns a timeTime object
// For a given year and month (as ints) it calculates the
// last date in the month.
func GetDateLastOfMonth(year, month int) (time.Time, error) {
	var date time.Time
	date, err := time.Parse(dateFormat, fmt.Sprintf(isoFormat, year, month, 1))
	if err != nil {
		return time.Time{}, err
	}
	date = date.AddDate(0, 1, -1)
	return date, nil
}
