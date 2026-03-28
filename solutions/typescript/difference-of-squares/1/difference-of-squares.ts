export class Squares {
  num: number;
  sumOfSquaresResult: number;
  squareOfSumResult: number;

  constructor(count: number) {
    this.num = count;
    this.sumOfSquaresResult = 0;
    this.squareOfSumResult = 0;
  }

  get sumOfSquares(): number {
    let sum = 0;
    for (let i = 1; i <= this.num; i++) {
      sum += i * i;
    }
    this.sumOfSquaresResult = sum;
    return this.sumOfSquaresResult;
  }

  get squareOfSum(): number {
    let sum = 0;
    for (let i = 1; i <= this.num; i++) {
      sum += i;
    }

    this.squareOfSumResult = sum * sum;
    return this.squareOfSumResult;
  }

  get difference(): number {
      return this.squareOfSumResult - this.sumOfSquaresResult;
  }
}
