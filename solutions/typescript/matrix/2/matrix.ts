export class Matrix {

  matrix: string
  mrows: number[][] = []
  mcolumns: number[][] = []

  constructor(matrix: string) {
      this.matrix = matrix

      for (let row of matrix.split('\n')) {
          this.mrows.push(row.split(' ').map(Number))
      }

      for (let i = 0; i < this.mrows[0].length; i++) {
          this.mcolumns.push(this.mrows.map(row => row[i]))
      }

      return this

  }

  get rows(): number[][] {
    return this.mrows
  }

  get columns(): number[][] {
    return this.mcolumns
  }
}
