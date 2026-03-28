export function hey(message: string): string {
    message = message.trim()

    let hasLetter : RegExpMatchArray | null = message.match(/[a-zA-Z]/)

    if (message.endsWith('?')) {
        if (hasLetter && message.toUpperCase() == message) {
            return 'Calm down, I know what I\'m doing!'
        } else {
            return 'Sure.'
        }
    }
    if (message.trim() == '') {
        return 'Fine. Be that way!'
    }
    if (hasLetter && message.toUpperCase() == message) {
        return 'Whoa, chill out!'
    }
    return 'Whatever.'
}
