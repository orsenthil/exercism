export function verse(verseNumber: number): string[] {
  if (verseNumber === 1) {
    return ['This is the house that Jack built.']
  }
  if (verseNumber === 2) {
    return [
      'This is the malt',
      'that lay in the house that Jack built.',
    ]
  }
  if (verseNumber === 3) {
    return [
      'This is the rat',
      'that ate the malt',
      'that lay in the house that Jack built.',
    ]
  }
  if (verseNumber === 4) {
    return [
      'This is the cat',
      'that killed the rat',
      'that ate the malt',
      'that lay in the house that Jack built.',
    ]
  }
  if (verseNumber === 5) {
    return [
      'This is the dog',
      'that worried the cat',
      'that killed the rat',
      'that ate the malt',
      'that lay in the house that Jack built.',
    ]
  }
  if (verseNumber === 6) {
    return [
      'This is the cow with the crumpled horn',
      'that tossed the dog',
      'that worried the cat',
      'that killed the rat',
      'that ate the malt',
      'that lay in the house that Jack built.',
    ]
  }
  if (verseNumber === 7) {
    return [
      'This is the maiden all forlorn',
      'that milked the cow with the crumpled horn',
      'that tossed the dog',
      'that worried the cat',
      'that killed the rat',
      'that ate the malt',
      'that lay in the house that Jack built.',
    ]
  }

  if (verseNumber === 8) {
    return  [
      'This is the man all tattered and torn',
      'that kissed the maiden all forlorn',
      'that milked the cow with the crumpled horn',
      'that tossed the dog',
      'that worried the cat',
      'that killed the rat',
      'that ate the malt',
      'that lay in the house that Jack built.',
    ]
  }

  if (verseNumber === 9) {
    return [
        'This is the priest all shaven and shorn',
        'that married the man all tattered and torn',
        'that kissed the maiden all forlorn',
        'that milked the cow with the crumpled horn',
        'that tossed the dog',
        'that worried the cat',
        'that killed the rat',
        'that ate the malt',
        'that lay in the house that Jack built.',
      ]
    }

    if (verseNumber === 10) {
      return [
        'This is the rooster that crowed in the morn',
        'that woke the priest all shaven and shorn',
        'that married the man all tattered and torn',
        'that kissed the maiden all forlorn',
        'that milked the cow with the crumpled horn',
        'that tossed the dog',
        'that worried the cat',
        'that killed the rat',
        'that ate the malt',
        'that lay in the house that Jack built.',
      ]
    }

    if (verseNumber === 11) {
      return  [
        'This is the farmer sowing his corn',
        'that kept the rooster that crowed in the morn',
        'that woke the priest all shaven and shorn',
        'that married the man all tattered and torn',
        'that kissed the maiden all forlorn',
        'that milked the cow with the crumpled horn',
        'that tossed the dog',
        'that worried the cat',
        'that killed the rat',
        'that ate the malt',
        'that lay in the house that Jack built.',
      ]
    }

    if (verseNumber === 12) {
      return [
        'This is the horse and the hound and the horn',
        'that belonged to the farmer sowing his corn',
        'that kept the rooster that crowed in the morn',
        'that woke the priest all shaven and shorn',
        'that married the man all tattered and torn',
        'that kissed the maiden all forlorn',
        'that milked the cow with the crumpled horn',
        'that tossed the dog',
        'that worried the cat',
        'that killed the rat',
        'that ate the malt',
        'that lay in the house that Jack built.',
      ]
    }

  return []
}

export function verses(startVerse: number, endVerse: number): string[] {
  let lyrics: string[] = []

  for (let i = startVerse; i <= endVerse; i++) {
    lyrics = lyrics.concat(verse(i))
    if (i < endVerse) {
      lyrics.push('')
    }
  }

    return lyrics
}
