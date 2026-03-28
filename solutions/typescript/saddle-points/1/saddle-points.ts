export function saddlePoints(grid: number[][]): { row: number, column: number }[] {
  let result: { row: number, column: number }[] = []
  
  for (let i = 0; i < grid.length; i++) {
    let row = grid[i]
    let max = Math.max(...row)
    let maxIndices: number[] = []
    row.forEach((value, index) => { if (value === max) maxIndices.push(index)})
    for (let maxIndex of maxIndices) {
      let column = grid.map(row => row[maxIndex])
      let min = Math.min(...column)
      if (max === min) {
        result.push({ row: i + 1, column: maxIndex + 1 })
      }
    }
  }

  return result
}
