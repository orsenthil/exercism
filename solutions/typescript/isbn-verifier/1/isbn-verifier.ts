export function isValid(isbn: string): boolean {
  if (isbn.match(/[^0-9X-]/)) return false
  let digits = isbn.replace(/-/g, '').split('')
  if (digits.length !== 10) return false
  if (digits[9] === 'X') digits[9] = '10'
  let sum = 0
  for (let i = 0; i < digits.length; i++) {
    sum += parseInt(digits[i]) * (10 - i)
  }
  if (sum % 11 === 0) return true
  return false
}
