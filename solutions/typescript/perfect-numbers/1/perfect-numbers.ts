export function classify(input: number): 'perfect' | 'abundant' | 'deficient' {
  if (input <= 0) {
    throw new Error('Classification is only possible for natural numbers.');
  }
  let factors = [];

  for (let i = 1; i < input; i++) {
    if (input % i === 0) {
      factors.push(i);
    }
  }
  if (factors.reduce((a, b) => a + b, 0) === input) {
    return 'perfect';
  }

  if (factors.reduce((a, b) => a + b, 0) > input) {
    return 'abundant';
  }

  return 'deficient';
}
