export function isPangram(sentence: string): boolean {
  var alphabet = 'abcdefghijklmnopqrstuvwxyz'
  sentence = sentence.toLowerCase()
  for (var i = 0; i < alphabet.length; i++) {
      if (! sentence.includes(alphabet[i])) {
          return false
      }
  }
  return true
}
