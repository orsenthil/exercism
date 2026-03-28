package armstrong

import (
	"math"
	"strconv"
)

// IsNumber checks if a number is an Armstrong number.
// An Armstrong number is a number that is the sum of its own digits each raised to the power of the number of digits.
// For example, 153 is an Armstrong number because 1^3 + 5^3 + 3^3 = 153.
func IsNumber(n int) bool {
	var result int

	ns := strconv.Itoa(n)
	ndigits := len(ns)

	for _, c := range ns {
		num, _ := strconv.Atoi(string(c))
		result += int(math.Pow(float64(num), float64(ndigits)))
	}

	return result == n
}
