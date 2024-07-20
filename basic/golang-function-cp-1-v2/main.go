package main

import (
	"fmt"
	"strconv"
)

func DateFormat(day, month, year int) string {
	var date string
	var finalDay string
	var finalMonth string
	var finalYear string

	if day < 10 {
		var newDay = strconv.Itoa(day)
		finalDay = "0" + newDay
	} else {
		finalDay = strconv.Itoa(day)
	}

	switch month {
	case 1:
		finalMonth = "January"
	case 2:
		finalMonth = "February"
	case 3:
		finalMonth = "March"
	case 4:
		finalMonth = "April"
	case 5:
		finalMonth = "May"
	case 6:
		finalMonth = "June"
	case 7:
		finalMonth = "July"
	case 8:
		finalMonth = "August"
	case 9:
		finalMonth = "September"
	case 10:
		finalMonth = "October"
	case 11:
		finalMonth = "November"
	case 12:
		finalMonth = "December"
	default:
		finalMonth = "Invalid month"
	}

	finalYear = strconv.Itoa(year)

	date = finalDay + "-" + finalMonth + "-" + finalYear

	return date
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(31, 12, 2020))
}
