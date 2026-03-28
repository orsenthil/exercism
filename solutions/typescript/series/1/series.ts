export class Series {
  series: string
  constructor(series: string) {
    this.series = series as string
  }

  slices(sliceLength: number): number[][] {
    if (this.series === '') {
      throw new Error("series cannot be empty")
    }
    const seriesLength = this.series.length
    if (sliceLength > seriesLength) {
      throw new Error("slice length cannot be greater than series length")
    }
    if (sliceLength < 0) {
      throw new Error("slice length cannot be negative") 
    }

    if (sliceLength === 0) {
      throw new Error("slice length cannot be zero")
    }
    const result: number[][] = []
    for (let i = 0; i < seriesLength - sliceLength + 1; i++) {
      let slice = this.series.slice(i, i + sliceLength)
      let sliceArr = slice.split('')
      let sliceNum = sliceArr.map(Number)
      result.push(sliceNum)
    }

    return result
  }
}
