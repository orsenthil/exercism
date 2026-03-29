package bafflingbirthdays

import "time"
import "math/rand"


func SharedBirthday(dates []time.Time) bool {
	seen := make(map[string]bool, len(dates))

	for _, date := range dates {
		if seen[date.Format("01-02")] {
			return true
		}
		seen[date.Format("01-02")] = true
	}
	return false
}

func RandomBirthdates(size int) []time.Time {
	dates := make([]time.Time, size)
	for i := range dates {
		dates[i] = time.Date(1981, time.Month(rand.Intn(12)+1), rand.Intn(31)+1, 0, 0, 0, 0, time.UTC)
	}
	return dates
}

func EstimatedProbability(size int) float64 {
	var trials int
	for range 1000 {
		dates := RandomBirthdates(size)
		if SharedBirthday(dates) {
			trials++
		}
	}
	return (float64(trials) / 1000.0) * 100.0
}