package sieve

func Sieve(limit int) []int {
	var primes []int
	var sieve = make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		if !sieve[i] {
			primes = append(primes, i)
			for j := i * i; j <= limit; j += i {
				sieve[j] = true
			}
		}
	}
	return primes

}
