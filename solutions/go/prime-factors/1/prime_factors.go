package prime

func Factors(n int64) []int64 {
	// prime factors of a number
	var result []int64

	// loop through all numbers from 2 to n
	for i := int64(2); i <= n; i++ {
		// if n is divisible by i
		for n%i == 0 {
			// add i to result
			result = append(result, i)
			// divide n by i
			n = n / i
		}
	}
	return result
}
