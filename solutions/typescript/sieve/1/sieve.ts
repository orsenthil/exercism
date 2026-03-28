export function primes(input: number): number[]{
  // Implement Sieve of Eratosthenes
  let primes: number[] = []
  let sieve: boolean[] = []
  for (let i = 2; i <= input; i++) {
    if (!sieve[i]) {
      primes.push(i)
      for (let j = i * 2; j <= input; j += i) {
        sieve[j] = true
      }
    }
  }

  return primes
}
