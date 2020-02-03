package main

import (
	"fmt"

	"github.com/jakoubek/dates"
)

func main() {
	first, _ := dates.GetDateFirstOfMonth(2019, 12)
	last, _ := dates.GetDateLastOfMonth(2019, 12)
	fmt.Println("First day of the month:", first.Format("2006-01-02"))
	fmt.Println("Last day of the month :", last.Format("2006-01-02"))
}
