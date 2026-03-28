export function verse(verseNumber: number): string {
  if (verseNumber === 1) {
    return `I know an old lady who swallowed a fly.
I don't know why she swallowed the fly. Perhaps she'll die.
`
  }
  if (verseNumber === 2) {
    return `I know an old lady who swallowed a spider.
It wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.
`
  }
  if (verseNumber === 3) {
    return `I know an old lady who swallowed a bird.
How absurd to swallow a bird!
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.
`
  }
  if (verseNumber === 4) {
    return `I know an old lady who swallowed a cat.
Imagine that, to swallow a cat!
She swallowed the cat to catch the bird.
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.
`
  }
  if (verseNumber === 5) {
    return `I know an old lady who swallowed a dog.
What a hog, to swallow a dog!
She swallowed the dog to catch the cat.
She swallowed the cat to catch the bird.
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.
`
  }
  if (verseNumber === 6) {
    return `I know an old lady who swallowed a goat.
Just opened her throat and swallowed a goat!
She swallowed the goat to catch the dog.
She swallowed the dog to catch the cat.
She swallowed the cat to catch the bird.
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.
`
  }
  if (verseNumber === 7) {
    return `I know an old lady who swallowed a cow.
I don't know how she swallowed a cow!
She swallowed the cow to catch the goat.
She swallowed the goat to catch the dog.
She swallowed the dog to catch the cat.
She swallowed the cat to catch the bird.
She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.
She swallowed the spider to catch the fly.
I don't know why she swallowed the fly. Perhaps she'll die.
`
  }
  if (verseNumber === 8) {
    return `I know an old lady who swallowed a horse.
She's dead, of course!
`
  }
  return ''
}

export function verses(startVerse: number, endVerse: number): string {
  let result = ''
  for (let i = startVerse; i <= endVerse; i++) {
    result += verse(i)
    if (i !== endVerse) {
      result += '\n'
    }
  }
  return result
}
