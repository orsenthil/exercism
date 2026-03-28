export function ofSize(input: number): number[][] {
  let matrix: number[][] = []
  let counter = 1
  let row = 0
  let col = 0
  let size = input
  let direction = 'right'

  for (let i = 0; i < size; i++) {
    matrix.push([])
    for (let j = 0; j < size; j++) {
      matrix[i].push(0)
    }
  }

  while (counter <= size * size) {
    matrix[row][col] = counter

    if (direction === 'right') {
      if (col + 1 < size && matrix[row][col + 1] === 0) {
        col++
      } else {
        direction = 'down'
        row++
      }
    } else if (direction === 'down') {
      if (row + 1 < size && matrix[row + 1][col] === 0) {
        row++
      } else {
        direction = 'left'
        col--
      }
    } else if (direction === 'left') {
      if (col - 1 >= 0 && matrix[row][col - 1] === 0) {
        col--
      } else {
        direction = 'up'
        row--
      }
    } else if (direction === 'up') {
      if (row - 1 >= 0 && matrix[row - 1][col] === 0) {
        row--
      } else {
        direction = 'right'
        col++
      }
    }

    counter++
  }

  return matrix
}
