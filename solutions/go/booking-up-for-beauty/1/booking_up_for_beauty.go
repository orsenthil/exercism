package booking

import (
    "fmt"
    "time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
    timeT, _ := time.Parse("1/2/2006 15:04:05", date)
	return timeT
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
    currentTime := time.Now()
    schedule, _ := time.Parse("January 2, 2006 15:04:05", date)
    return schedule.Before(currentTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
    timeT, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)

    if (timeT.Hour() >= 12 && timeT.Hour() <= 18) {
        return true
    }

    return false
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
    timeT := Schedule(date)
    return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", timeT.Weekday(), timeT.Month(), timeT.Day(), timeT.Year(), timeT.Hour(), timeT.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
    currentTime := time.Now()
    currentTimeStr := fmt.Sprintf("%d-09-15", currentTime.Year())
    anniversary, _ := time.Parse("2006-01-02", currentTimeStr)
	
    return anniversary
}
