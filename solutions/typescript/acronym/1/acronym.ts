export function parse(phrase: string): string {
    const words = phrase.split(/[\s-]/)
    return words.map(word => word[0].toUpperCase()).join('')
}
