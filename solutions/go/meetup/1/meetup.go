package meetup

import "time"

// Define the WeekSchedule type here.

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	// get the first day of the month
	firstDayOfTheMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	lastDayOfTheMonth := firstDayOfTheMonth.AddDate(0, 1, -1)

	currentDate := firstDayOfTheMonth

	for currentDate.Before(lastDayOfTheMonth) {
		if currentDate.Weekday() == wDay {
			switch wSched {
			case First:
				return currentDate.Day()
			case Second:
				return currentDate.Day() + 7
			case Third:
				return currentDate.Day() + 14
			case Fourth:
				return currentDate.Day() + 21
			case Teenth:
				if currentDate.Day() >= 13 && currentDate.Day() <= 19 {
					return currentDate.Day()
				}
			case Last:
				if currentDate.Day() >= 25 {
					return currentDate.Day()
				}

				if currentDate.AddDate(0, 0, 7).Month() != month {
					return currentDate.Day()
				}
			}
		}
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	if currentDate.Equal(lastDayOfTheMonth) && currentDate.Weekday() == wDay {
		return currentDate.Day()
	}

	return -1
}
