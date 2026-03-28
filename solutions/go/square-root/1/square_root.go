package squareroot

import "errors"

func SquareRoot(number int) (int, error) {
	if number < 0 {
		return 0, errors.New("number is negative")
	}
	var guess int = 1
	for guess*guess != number {
		guess = (guess + number / guess) / 2
	}
	return guess, nil
}
