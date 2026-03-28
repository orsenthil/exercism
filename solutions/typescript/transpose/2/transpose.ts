export function transpose(input: string[]): string[] {
  let result: string[] = []
  let max = 0

  if (input.length === 0) {
    return []
  }

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


  result[result.length - 1] = result[result.length - 1].trimEnd()
  return result

}
