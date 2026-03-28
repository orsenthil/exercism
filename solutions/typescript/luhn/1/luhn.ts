export function valid(digitString: string): Boolean {
  digitString = digitString.replace(/\s/g, '')
  if (digitString.length <= 1) return false

  let sum = 0
  let isEven = digitString.length % 2 === 0

  for (let i = 0; i < digitString.length; i++) {
    let digit = parseInt(digitString[i])
    if (isNaN(digit)) return false
    if (isEven) {
      if (i % 2 === 0) {
        digit *= 2
        if (digit > 9) digit -= 9
      }
    } else {
      if (i % 2 !== 0) {
        digit *= 2
        if (digit > 9) digit -= 9
      }
    }
    sum += digit
  }

  return sum % 10 === 0
}
