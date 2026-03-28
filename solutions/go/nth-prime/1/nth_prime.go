package prime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n must be greater than 1")

	}

	count := 1

	for i := 2; i <= n*n; i++ {
		if isPrime(i) {
			if count == n {
				return i, nil
			}
			count++
		}
	}
	return 2, nil
}

// isPrime returns true if the number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
