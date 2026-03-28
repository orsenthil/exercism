package summultiples

func SumMultiples(limit int, divisors ...int) int {
	var total int
	for i := 1; i < limit; i++ {
		for _, divisor := range divisors {
			if divisor != 0 && i%divisor == 0 {
				total += i
				break
			}
		}
	}
	return total
}
