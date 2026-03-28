package clock

import (
	"fmt"
	"strconv"
)

type Clock struct {
	hours   int
	minutes int
}

func New(h, m int) Clock {
	if h >= 0 && m >= 0 {
		if m >= 60 {
			h = h + (m / 60)
			m = m % 60
		}
		if h >= 24 {
			h = h % 24
		}

		return Clock{h, m}
	}
	if h >= 0 && m < 0 {
		h = h - (m / 60)
		h = h % 24
		if h < 0 {
			h += 24
		}
		m = m % 60
		if m < 0 {
			m += 60
		}

		return Clock{h, m}
	}
	m = m % 60
	if m < 0 {
		m += 60
	}
	h = h % 24
	if h < 0 {
		h += 24
	}

	if m >= 60 {
		h = h + (m / 60)
		m = m % 60
	}

	if h >= 24 {
		h = h % 24
	}

	return Clock{h, m}
}

func (c Clock) Add(m int) Clock {
	hours := c.hours
	minutes := c.minutes

	if minutes+m >= 60 {
		minutes = (minutes + m) % 60
		hours = (hours + 1) % 24
	} else {
		minutes = minutes + m
	}

	return Clock{hours, minutes}
}

func (c Clock) Subtract(m int) Clock {
	hours := c.hours
	minutes := c.minutes

	if minutes-m <= 0 {
		minutes = 60 - (minutes - m)
		hours = hours - 1
		if hours < 0 {
			hours = 24 - hours
		}
	} else {
		minutes = minutes - m
	}

	return Clock{hours, minutes}

}

func (c Clock) String() string {
	hours := strconv.Itoa(c.hours)
	minutes := strconv.Itoa(c.minutes)
	return fmt.Sprintf("%02s:%02s", hours, minutes)
}
