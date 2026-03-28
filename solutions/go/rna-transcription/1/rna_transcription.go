package strand

func ToRNA(dna string) string {
	var rna []rune
	for _, c := range dna {
		if c == 'G' {
			rna = append(rna, 'C')
		} else if c == 'C' {
			rna = append(rna, 'G')
		} else if c == 'T' {
			rna = append(rna, 'A')
		} else if c == 'A' {
			rna = append(rna, 'U')
		}
	}
	return string(rna)
}
