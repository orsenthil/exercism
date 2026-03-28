export const square = (cell: number): bigint => {
  if (cell < 1 || cell > 64) {
    throw new Error('square must be between 1 and 64')
  }
  return BigInt(2 ** (cell - 1))
}

export const total = (): bigint => {
  let sum: bigint = BigInt(0)
  for (let i = 1; i <= 64; i++) {
    sum += square(i);
  }
  return sum
}
