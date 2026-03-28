package protein

import "errors"

var ErrStop error = errors.New("stop")
var ErrInvalidBase error = errors.New("invalid base")

var mapping = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	// break the rna string into codons.
	// a group of 3 characters is a codon.

	var response []string

	for i := 0; i < len(rna); i++ {
		if i%3 == 0 {
			codon := rna[i : i+3]
			// if codon is a stop codon, return the proteins
			// if codon is not a stop codon, add the protein to the proteins slice
			acid, err := FromCodon(codon)
			if err == ErrInvalidBase {
				return response, err
			} else if err == ErrStop {
				return response, nil
			}

			response = append(response, acid)
		}

	}

	return response, nil
}

func FromCodon(codon string) (string, error) {
	/**
	Codon                 | Protein
	:---                  | :---
	AUG                   | Methionine
	UUU, UUC              | Phenylalanine
	UUA, UUG              | Leucine
	UCU, UCC, UCA, UCG    | Serine
	UAU, UAC              | Tyrosine
	UGU, UGC              | Cysteine
	UGG                   | Tryptophan
	UAA, UAG, UGA         | STOP
	*/

	acid, ok := mapping[codon]

	if !ok {
		return "", ErrInvalidBase
	} else if acid == "STOP" {
		return "", ErrStop
	}

	return acid, nil
}
