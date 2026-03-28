export class Matrix {
  matrix: string
  mrows: number[][]
  mcolumns: number[][]

  constructor(matrix: string) {
    this.matrix = matrix
    this.mrows = matrix.split('\n').map(row => row.split(' ').map(Number))
    this.mcolumns = this.mrows[0].map((_, i) => this.mrows.map(row => row[i]))

    return this

  }

  get rows(): unknown {
    return this.mrows
  }

  get columns(): unknown {
    return this.mcolumns
  }
}