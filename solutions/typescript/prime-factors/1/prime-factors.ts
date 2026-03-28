export function calculatePrimeFactors(input: number): number[] { 
  let factors: number[] = [];
  let divisor = 2;

  if (input === 1) {
    return factors;
  }

  if (input === 2) {
    factors.push(input);
    return factors;
  }

  while (input >= 2) {
    if (input % divisor === 0) {
      factors.push(divisor);
      input = input / divisor;
    } else {
      divisor++;
    }
  }
  return factors;
}
