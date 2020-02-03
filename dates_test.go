package dates

import (
	"fmt"
	"testing"
	"time"
)

func TestGetDateFirstOfMonth(t *testing.T) {
	var cases = []struct {
		year  int
		month int
		want  string
	}{
		{2020, 1, "2020-01-01"},
		{2020, 12, "2020-12-01"},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%04d-%02d", c.year, c.month), func(t *testing.T) {
			in, _ := GetDateFirstOfMonth(c.year, c.month)
			want, _ := time.Parse("2006-01-02", c.want)
			if !in.Equal(want) {
				t.Errorf("got %s, want %s", in.Format(dateFormat), want.Format(dateFormat))
			}
		})
	}
}

func TestGetDateLastOfMonth(t *testing.T) {
	var cases = []struct {
		year  int
		month int
		want  string
	}{
		{2019, 1, "2019-01-31"},
		{2019, 2, "2019-02-28"},
		{2020, 2, "2020-02-29"},
		{2020, 3, "2020-03-31"},
		{2020, 11, "2020-11-30"},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%04d-%02d", c.year, c.month), func(t *testing.T) {
			in, _ := GetDateLastOfMonth(c.year, c.month)
			want, _ := time.Parse("2006-01-02", c.want)
			if !in.Equal(want) {
				t.Errorf("got %s, want %s", in.Format(dateFormat), want.Format(dateFormat))
			}
		})
	}
}
