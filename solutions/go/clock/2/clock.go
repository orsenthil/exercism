package clock

import "fmt"

type Clock struct {
	hours   int
	minutes int
}

func constrain(h, m int) (int, int) {
	for m > 59 {
		h += 1
		m -= 60
	}

	for m < 0 {
		h -= 1
		m += 60
	}

	for h >= 24 {
		h -= 24
	}

	for h < 0 {
		h += 24
	}

	return h, m

}

func New(h, m int) Clock {

	h, m = constrain(h, m)

	return Clock{h, m}

}

func (c Clock) Add(m int) Clock {
	return New(c.hours, c.minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hours, c.minutes-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
