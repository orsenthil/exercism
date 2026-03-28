export const largestProduct = (inputString: string, span: number) => {

  if (span < 0) {
    throw new Error('Span must not be negative')
  }

  if (span > inputString.length) {
    throw new Error('Span must be smaller than string length')
  }

  if (inputString.match(/\D/)) {
    throw new Error('Digits input must only contain digits')
  }

  if (span == inputString.length) {
    return inputString.split('').reduce((acc, digit) => acc * parseInt(digit), 1)
  }
  let maxProduct = 0
  for (let i = 0; i <= inputString.length - span; i++) {
    let product = inputString.slice(i, i + span).split('').reduce((acc, digit) => acc * parseInt(digit), 1)
    if (product > maxProduct) {
      maxProduct = product
    }
  }

  return maxProduct
}
