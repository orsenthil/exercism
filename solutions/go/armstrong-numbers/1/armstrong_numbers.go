package armstrong

import (
	"math"
	"strconv"
)

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
