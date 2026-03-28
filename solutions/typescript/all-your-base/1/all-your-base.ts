export function convert(
  digits: number[],
  inputBase: number,
  outputBase: number 
): number[]{
  switch (true) {
    case inputBase <= 1: throw Error("Wrong input base")
    case outputBase <= 1: throw Error("Wrong output base")
    case outputBase % 1 !== 0: throw Error("Wrong output base")
    case inputBase % 1 !== 0: throw Error("Wrong input base")
    case digits.length === 0: throw Error("Input has wrong format")
    case digits.some(digit => digit < 0): throw Error("Input has wrong format")
    case digits.some(digit => digit >= inputBase): throw Error("Input has wrong format")
    case digits.length > 1 && digits[0] === 0: throw Error("Input has wrong format")
  }

  return toBase(fromBase(digits, inputBase), outputBase)

}

function fromBase(digits: number[], base: number): number {
  return digits.reverse().reduce((acc, digit, index) => {
    return acc + digit * base ** index
  }, 0)

}

function toBase(number: number, base: number): number[] {
  if (number === 0)  {
    return [0]
  }
  let output = []
  while (number > 0) {
    output.push(number % base)
    number = Math.floor(number / base)
  }

  return output.reverse()
}