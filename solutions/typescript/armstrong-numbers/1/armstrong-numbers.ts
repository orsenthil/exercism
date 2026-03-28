export function isArmstrongNumber(number: number | BigInt): boolean {
  if (typeof number === 'number') {
    if (number < 0) {
      return false
    }
    const numStr = number.toString()
    const numDigits = numStr.length
    return number === numStr.split('').reduce((acc, digit) => acc + Number(digit) ** numDigits, 0)
  } else {
    const numStr = number.toString()
    const numDigits = BigInt(numStr.length)
    return number === numStr.split('').reduce((acc, digit) => acc + BigInt(digit) ** numDigits, 0n)
  }
}
