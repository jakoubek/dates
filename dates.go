package dates

import (
	"fmt"
	"time"
)

var (
	dateFormat = "2006-01-02"
	isoFormat  = "%04d-%02d-%02d"
)

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
