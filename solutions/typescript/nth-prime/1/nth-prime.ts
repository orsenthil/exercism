export function nth(input: number): number {

  if (input < 1) {
    throw new Error('Prime is not possible');
  }

  let primes = [2];
  let i = 3;

  while (primes.length < input) {
    if (primes.every(prime => i % prime !== 0)) {
      primes.push(i);
    }
    i += 2;
  }

  return primes[primes.length - 1];

}
