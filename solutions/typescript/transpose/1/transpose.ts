export function transpose(input: string[]): string[] {
  let result: string[] = []
  let max = 0
  input.forEach((line) => {
    if (line.length > max) {
      max = line.length
    }
  })
  
  for (let i = 0; i < max; i++) {
    let line = ''
    input.forEach((row) => {
      if (i < row.length) {
        line += row[i]
      } else {
        line += ' '
      }
    })
    result.push(line)
  }

  for (let i = 0; i < result.length; i++) {
    result[i] = result[i].trimEnd()
  }

  return result

}
