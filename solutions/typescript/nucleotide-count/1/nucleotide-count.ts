type Nucleotide = 'A' | 'C' | 'G' | 'T';
export function nucleotideCounts(nucleotides: string) {
  let nucleotideCounts = {
    'A': 0,
    'C': 0,
    'G': 0,
    'T': 0,
  }
  // handle invalid nucleotides
  if (nucleotides.match(/[^ACGT]/)) {
    throw new Error('Invalid nucleotide in strand')
  }

  for (const nucleotide of nucleotides) {
    nucleotideCounts[nucleotide as Nucleotide]++
  }

   return nucleotideCounts
}
