package algorithms

import "fmt"

func dayOfProgrammer(year int32) string {
	var programmerDayDate string
	var isLeap bool
	isLeap = false

	// which calendar, 1918 special case. year of change of calendar
	if year == 1918 {
		programmerDayDate = "26.09.1918"

	} else if year < 1918 {
		// julian calendar
		// in julian calendar, leap year is a year divisible by 4
		if year%4 == 0 {
			isLeap = true
		} else {
			isLeap = false
		}

		if isLeap == true {
			programmerDayDate = fmt.Sprintf("12.09.%d", year)
		} else {
			programmerDayDate = fmt.Sprintf("13.09.%d", year)
		}

	} else {
		// gregorian calendar

		// in gregorian calendar, leap year if divisible by 400 or divisible by 4 and not by 100
		// what even is a leap year
		if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
			isLeap = true
		} else {
			isLeap = false
		}

		if isLeap == true {
			programmerDayDate = fmt.Sprintf("12.09.%d", year)
		} else {
			programmerDayDate = fmt.Sprintf("13.09.%d", year)
		}
	}
	return programmerDayDate
}
