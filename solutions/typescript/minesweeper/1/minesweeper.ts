export function annotate(field: Array<string>): Array<string> {
  let result: Array<string> = []
  for (let i = 0; i < field.length; i++) {
    result.push('')
    for (let j = 0; j < field[i].length; j++) {
      let numBomb = 0;
      if (field[i][j] !== '*') {
        for (let h = -1; h <= 1; h++) {
          for (let w = -1; w <= 1; w++) {
            if ((i + h >= 0) && (i + h < field.length) && (j + w >= 0) && (j + w < field[i].length)) {
              if (field[i + h][j + w] === '*') {
                numBomb++
              }
            }
          }
        }
        if (numBomb > 0) {
          result[i] += numBomb.toString()
          numBomb = 0
        } else {
          result[i] += ' '
        }
      } else {
        result[i] += '*'
      }
    }
  }
  return result;
}
