#include "protein_translation.h"
#include <string.h>

protein_t codon_to_polypeptide(const char *codon);

protein_t codon_to_polypeptide(const char *codon) {
    if (strcmp(codon, "AUG") == 0) return Methionine;
    else if (strcmp(codon, "UUU") == 0 || strcmp(codon, "UUC") == 0) return Phenylalanine;
    else if (strcmp(codon, "UUA") == 0 || strcmp(codon, "UUG") == 0) return Leucine;
    else if (strcmp(codon, "UCU") == 0 || strcmp(codon, "UCC") == 0 || strcmp(codon, "UCA") == 0 || strcmp(codon, "UCG") == 0) return Serine;
    else if (strcmp(codon, "UAU") == 0 || strcmp(codon, "UAC") == 0) return Tyrosine;
    else if (strcmp(codon, "UGU") == 0 || strcmp(codon, "UGC") == 0) return Cysteine;
    else if (strcmp(codon, "UGG") == 0) return Tryptophan;
    else if (strcmp(codon, "UAA") == 0 || strcmp(codon, "UAG") == 0 || strcmp(codon, "UGA") == 0) return Stop;
    else return Unknown;
}

proteins_t proteins(const char *const rna) {
    proteins_t protein;

    if (rna == NULL) {
        protein.valid = false;
        return protein;
    }

    protein.valid = true;
    protein.count = 0;

    size_t size = strlen(rna);
    size_t idx = 0;
    char codon[4];

    while (1) {
        if (protein.count == MAX_PROTEINS) {
            return protein;
        }

        if (idx + 3 > size) {
            protein.valid = idx == size;
            return protein;
        }
        strncpy(codon, rna + idx, 3);
        protein.proteins[protein.count] = codon_to_polypeptide(codon);
        if (protein.proteins[protein.count] == Stop || protein.proteins[protein.count] == Unknown) {
            protein.valid = protein.proteins[protein.count] == Stop;
            return protein;
        }

        protein.count++;
        idx += 3;
    }
}
