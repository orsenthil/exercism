interface Input {
  maxFactor: number
  minFactor?: number
}

interface Palindrome {
  value: number | null
  factors: number[][]
}

interface Output {
  smallest: Palindrome
  largest: Palindrome
}

export function generate(params: Input): Output {
  let {minFactor, maxFactor} = params
  if (!minFactor || !maxFactor) throw new Error("Must provide 'minFactor' and 'maxFactor' in params object")
  if (minFactor > maxFactor) throw new Error("min must be <= max")
  
  let smallest: Palindrome = {value: null, factors: []}
  let largest: Palindrome = {value: null, factors: []}

  for (let i = minFactor; i <= maxFactor; i++) {
    for (let j = i; j <= maxFactor; j++) {
      let prod = i * j

      if (isPalindrome(prod)) {
        if (!smallest.value || prod < smallest.value) {
          smallest.value = prod 
          smallest.factors = [[i, j]]
        } else if (prod === smallest.value) {
          smallest.factors.push([i, j])
        } else if (!largest.value || prod > largest.value) {
          largest.value = prod
          largest.factors = [[i, j]]
        } else if (prod === largest.value) {
          largest.factors.push([i, j])
        }
      }
    }
  }
  return {smallest, largest}
}

function isPalindrome(num: number): boolean {
  return num.toString() === num.toString().split("").reverse().join("")
}
  