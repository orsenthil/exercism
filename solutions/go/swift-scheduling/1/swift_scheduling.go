package swiftscheduling

import (
	"strconv"
	"strings"
	"time"
)

const layout = "2006-01-02T15:04:05"

func DeliveryDate(start, delivery string) string {
	t := parseTime(start)
	switch delivery {
	case "NOW":
		t = t.Add(2 * time.Hour)
	case "ASAP":
		if t.Hour() < 13 {
			t = time.Date(t.Year(), t.Month(), t.Day(), 17, 0, 0, 0, t.Location())
		} else {
			t = time.Date(t.Year(), t.Month(), t.Day()+1, 13, 0, 0, 0, t.Location())
		}
	case "EOW":
		switch t.Weekday() {
		case time.Monday, time.Tuesday, time.Wednesday:
			count := 5 - int(t.Weekday())
			t = time.Date(t.Year(), t.Month(), t.Day()+count, 17, 0, 0, 0, t.Location())
		case time.Thursday, time.Friday:
			count := 7 - int(t.Weekday())
			t = time.Date(t.Year(), t.Month(), t.Day()+count, 20, 0, 0, 0, t.Location())
		}
	default:
		if strings.HasSuffix(delivery, "M") {
			N, _ := strconv.Atoi(strings.TrimSuffix(delivery, "M"))
			year := t.Year()
			if int(t.Month()) >= N {
				year++
			}
			d := time.Date(year, time.Month(N), 1, 8, 0, 0, 0, t.Location())
			for d.Weekday() == time.Saturday || d.Weekday() == time.Sunday {
				d = d.AddDate(0, 0, 1)
			}
			t = d
		} else if strings.HasPrefix(delivery, "Q") {
			N, _ := strconv.Atoi(strings.TrimPrefix(delivery, "Q"))
			year := t.Year()
			if int(t.Month()) > N * 3 {
				year++
			}
			d := time.Date(year, time.Month(N*3 + 1), 0, 8, 0, 0, 0, t.Location())
			for d.Weekday() == time.Saturday || d.Weekday() == time.Sunday {
				d = d.AddDate(0, 0, -1)
			}
			t = d
		}
	}

	return formatTime(t)
}

func parseTime(t string) time.Time {
	parsedTime, err := time.Parse(layout, t)
	if err != nil {
		panic(err)
	}
	return parsedTime
}

func formatTime(t time.Time) string {
	return t.Format(layout)
}
