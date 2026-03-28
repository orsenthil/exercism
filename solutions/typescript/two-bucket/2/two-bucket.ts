function gcd(a: number, b: number): number {
  if (b === 0) {
    return a
  }
  return gcd(b, a % b)
}

export class TwoBucket {
  private _goalBucket = 'one'
  private _otherBucket = 0
  private _goal: number
  private _steps = 0
  private _starterBuck: boolean
  private _buckOne: number
  private _buckTwo: number

  constructor(buckOne: number, buckTwo: number, goal: number, starterBuck: string) {
    this._goal = goal
    this._starterBuck = starterBuck === 'one'
    this._buckOne = this._starterBuck ? buckOne : buckTwo
    this._buckTwo = this._starterBuck ? buckTwo : buckOne
  }

  moves() : number {
    if (this._goal % gcd(this._buckOne, this._buckTwo) !== 0) {
      throw new Error("Not possible to reach the goal ")
    }
    if (this._goal > this._buckOne && this._goal > this._buckTwo) {
      throw new Error("Not possible to reach the goal ")
    }

    if (this._goal === this._buckTwo) {
      this._goalBucket = 'two'
      this._otherBucket = this._buckOne
      return 2
    }

    let a = this._buckOne
    let b = 0

    this._steps += 1

    while (a !== this._goal && b !== this._goal) {
      const amt = Math.min(a, this._buckTwo - b)
      a -= amt
      b += amt
      this._steps += 1

      if (a === this._goal || b === this._goal) {
        break
      }

      if (a === 0) {
        a = this._buckOne
        this._steps += 1
      }

      if (b === this._buckTwo) {
        b = 0
        this._steps += 1
      }
    }

    this._goalBucket = a === this._goal 
    ? this._starterBuck ? 'one' : 'two'
    : this._starterBuck ? 'two' : 'one'

    this._otherBucket = a === this._goal ? b : a

    return this._steps
  }

  get goalBucket() {
    return this._goalBucket
  }

  get otherBucket() {
    return this._otherBucket
  }

}
