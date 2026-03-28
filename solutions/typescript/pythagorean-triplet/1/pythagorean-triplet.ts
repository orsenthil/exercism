type Options = {
  minFactor?: number
  maxFactor?: number
  sum: number
}

export function triplets(options: Options): Triplet[] {
  let minFactor = options.minFactor || 1
  let maxFactor = options.maxFactor || options.sum
  let triplets: Triplet[] = []
  for (let a = minFactor; a <= maxFactor; a++) {
    for (let b = a; b <= maxFactor; b++) {
      let c = options.sum - a - b
      if (c > b  && a * a + b * b === c * c) {
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
    return this.sides.sort((a, b) => a - b)
  }

}
