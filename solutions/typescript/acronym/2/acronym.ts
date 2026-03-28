export function parse(phrase: string): string {
    // replace camelCase with space
    phrase = phrase.replace(/([a-z])([A-Z])/g, '$1 $2')
    const words = phrase.split(/[\s-]/)
    return words.map(word => word[0].toUpperCase()).join('')
}
