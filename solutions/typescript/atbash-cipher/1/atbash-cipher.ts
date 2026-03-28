export function encode(plainText: string): string {
  const alphabet = 'abcdefghijklmnopqrstuvwxyz'
  for (let i = 0; i < plainText.length; i++) {
      const char = plainText[i]
      if (char.match(/[a-z]/i)) {
          const index = alphabet.indexOf(char.toLowerCase())
          const cipherChar = alphabet[25 - index]
          plainText = plainText.slice(0, i) + cipherChar + plainText.slice(i + 1)
      }
  }
  // remove non alphanumeric characters
  plainText = plainText.replace(/[^a-z0-9]/gi, '')
  // Group the cipher text in blocks of 5 characters
  for (let i = 5; i < plainText.length; i += 6) {
      plainText = plainText.slice(0, i) + ' ' + plainText.slice(i)
  }

  return plainText
}

export function decode(cipherText: string): string {
  const alphabet = 'abcdefghijklmnopqrstuvwxyz'
  for (let i = 0; i < cipherText.length; i++) {
      const char = cipherText[i]
      if (char.match(/[a-z]/i)) {
          const index = alphabet.indexOf(char.toLowerCase())
          const plainChar = alphabet[25 - index]
          cipherText = cipherText.slice(0, i) + plainChar + cipherText.slice(i + 1)
      }
  }

  // remove non alphanumeric characters
  cipherText = cipherText.replace(/[^a-z0-9]/gi, '')
  return cipherText
}
