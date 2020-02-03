# dates

dates is a helper package to handle dates in Golang.

## Installation

You can install this package in your project by running:

```
go get -u github.com/jakoubek/dates
```

## Available functions

- **GetDateFirstOfMonth(year, month int) (time.Time, error)** - Get the first day of a month. Year and month are ints.
- **GetDateLastOfMonth(year, month int) (time.Time, error)** - Get the last day of a month. Year and month are ints.
