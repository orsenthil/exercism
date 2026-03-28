export function translate(input: string): string {
  let result = ''
  let words = input.split(' ')
  for (let word of words) {
    let firstVowelIndex = word.search(/[aeiou]/)
    let firstLetter = word[0]
    if (firstVowelIndex === 0) {
      result += word + 'ay'
    } else if (firstLetter === 'q' && word[1] === 'u') {
      result += word.slice(2) + firstLetter + word[1] + 'ay'
    } else if (word[1] === 'q' && word[2] === 'u') {
      result += word.slice(3) + firstLetter + word[1] + word[2] + 'ay'
    } else {
      result += word.slice(firstVowelIndex) + word.slice(0, firstVowelIndex) + 'ay'
    }
    result += ' '
  }
  return result.trim()
}
