export function translate(protein: string) {
  let mapprotein = {
    'AUG': 'Methionine',
    'UUU': 'Phenylalanine',
    'UUC': 'Phenylalanine',
    'UUA': 'Leucine',
    'UUG': 'Leucine',
    'UCU': 'Serine',
    'UCC': 'Serine',
    'UCA': 'Serine',
    'UCG': 'Serine',
    'UAU': 'Tyrosine',
    'UAC': 'Tyrosine',
    'UGU': 'Cysteine',
    'UGC': 'Cysteine',
    'UGG': 'Tryptophan',
    'UAA': 'STOP',
    'UAG': 'STOP',
    'UGA': 'STOP'
  }

  let result = []

  for (let i = 0; i < protein.length; i += 3) {
    let codon = protein.slice(i, i + 3)
    let value = mapprotein[codon as keyof typeof mapprotein]

    if (value === undefined) {
      throw new Error('Invalid codon')
    }

    if (value === 'STOP') {
      break
    }

    result.push(value)

  }

  return result
}
