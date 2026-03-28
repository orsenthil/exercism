type Options = {
  minFactor?: number
  maxFactor?: number
  sum: number
}

export function triplets(options: Options): Triplet[] {
  let minFactor = options.minFactor || 1
  let maxFactor = options.maxFactor || Math.ceil(options.sum / 2 - 1)
  let triplets: Triplet[] = []
  for (let a = maxFactor; a > minFactor + 1; a--) {
    for (let b = a - 1; b > minFactor + 1; b--) {
      const c = options.sum - a - b
      if (c <= maxFactor && c >= minFactor && a * a + b * b === c * c) {
        triplets.push(new Triplet(a, b, c))
      }
    }
  }
  return triplets
}

class Triplet {
  private sides: [number, number, number]
  constructor(a: number, b: number, c: number) {
    this.sides = [a, b, c]
  }

  toArray(): [number, number, number] {
    return this.sides
  }
}
