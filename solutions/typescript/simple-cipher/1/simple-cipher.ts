export class SimpleCipher {
  key: string

  constructor(key? : string) {
    let randomKey = ''
    for (let i = 0; i < 100; i++) {
      randomKey += String.fromCharCode(Math.floor(Math.random() * 26) + 97)
    }
    this.key = key || randomKey
  }

  encode(message: string): string {
    let encodedMessage = ''
    for (let i = 0; i < message.length; i++) {
        let charCode = message.charCodeAt(i)
        let keyChar = this.key.charCodeAt(i % this.key.length)
        // add the keyChar to the charCode and subtract 97 because the charCode of 'a' is 97
        let newCharCode = charCode + keyChar - 97
        // if the newCharCode is greater than 122, which is the charCode of 'z', subtract 26 to wrap around
        if (newCharCode > 122) {
            newCharCode -= 26
        }
        encodedMessage += String.fromCharCode(newCharCode)
    }
    return encodedMessage
  }

  decode(message: string): string {
      let decodedMessage = ''
      for (let i = 0; i < message.length; i++) {
          let charCode = message.charCodeAt(i)
          let keyChar = this.key.charCodeAt(i % this.key.length)
          // subtract the keyChar from the charCode and add 97 because the charCode of 'a' is 97
          let newCharCode = charCode - keyChar + 97
          // if the newCharCode is less than 97, which is the charCode of 'a', add 26 to wrap around
          if (newCharCode < 97) {
              newCharCode += 26
          }
          decodedMessage += String.fromCharCode(newCharCode)
      }
      return decodedMessage
  }
}
