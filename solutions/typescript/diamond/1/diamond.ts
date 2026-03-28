export function makeDiamond(character: string): string {
  const alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
  const index = alphabet.indexOf(character)
  const diamond: string[] = []

  for (let i = 0; i <= index; i++) {
    const spaces = ' '.repeat(index - i)
    const letter = alphabet[i]
    if (i === 0) {
      diamond.push(`${spaces}${letter}${spaces}`)
    } else {
      const innerSpaces = ' '.repeat(2 * i - 1)
      diamond.push(`${spaces}${letter}${innerSpaces}${letter}${spaces}`)
    }
  }

  for (let i = index - 1; i >= 0; i--) {
    const spaces = ' '.repeat(index - i)
    const letter = alphabet[i]
    if (i === 0) {
      diamond.push(`${spaces}${letter}${spaces}`)
    } else {
      const innerSpaces = ' '.repeat(2 * i - 1)
      diamond.push(`${spaces}${letter}${innerSpaces}${letter}${spaces}`)
    }
  }

  return diamond.join('\n') + '\n'
}
