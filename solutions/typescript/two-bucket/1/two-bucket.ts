export class TwoBucket {
  bucketOne: number
  bucketTwo: number
  goal: number
  starterBucket: string
  finalgoalBucket: string
  finalOtherBucket: number

  constructor(buckOne: number, buckTwo : number, goal: number, starterBuck: string) {
    this.bucketOne = buckOne
    this.bucketTwo = buckTwo
    this.goal = goal
    this.starterBucket = starterBuck
    this.finalgoalBucket = ''
    this.finalOtherBucket = 0
  }

  moves() {
    while (this.bucketOne !== this.goal && this.bucketTwo !== this.goal) {
      if (this.starterBucket === 'one') {
        if (this.bucketOne === 0) {
          this.bucketOne = this.bucketOne + this.bucketTwo
          this.bucketTwo = 0
        } else {
          this.bucketTwo = this.bucketTwo + this.bucketOne
          this.bucketOne = 0
        }
      } else {
        if (this.bucketTwo === 0) {
          this.bucketTwo = this.bucketTwo + this.bucketOne
          this.bucketOne = 0
        } else {
          this.bucketOne = this.bucketOne + this.bucketTwo
          this.bucketTwo = 0
        }
      }
    }
    if (this.bucketOne === this.goal) {
      this.finalgoalBucket = 'one'
      return this.bucketTwo
    }
    this.finalgoalBucket = 'two'
    return this.bucketOne
  }

  get goalBucket() {
    return this.finalgoalBucket
  }

  get otherBucket() {
    if (this.finalgoalBucket === 'one') {
      return this.bucketTwo
    }
    return this.bucketOne
  }
}
