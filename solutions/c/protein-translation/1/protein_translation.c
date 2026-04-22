#include "protein_translation.h"
#include <string.h>
#include <stdlib.h>

amino_acid_t codon_to_amino_acid(const char *const codon) {
    if (codon == NULL || strlen(codon) != 3) {
        return 8; // Invalid codon
    }  

    if (strcmp(codon, "AUG") == 0) {
        return Methionine;
    } else if (strcmp(codon, "UUU") == 0 || strcmp(codon, "UUC") == 0) {
        return Phenylalanine;
    } else if (strcmp(codon, "UUA") == 0 || strcmp(codon, "UUG") == 0) {
        return Leucine;
    } else if (strcmp(codon, "UCU") == 0 || strcmp(codon, "UCC") == 0 ||
                strcmp(codon, "UCA") == 0 || strcmp(codon, "UCG") == 0) {
        return Serine;
    } else if (strcmp(codon, "UAU") == 0 || strcmp(codon, "UAC") == 0) {
        return Tyrosine;
    } else if (strcmp(codon, "UGU") == 0 || strcmp(codon, "UGC") == 0) {
        return Cysteine;
    } else if (strcmp(codon, "UGG") == 0) {
        return Tryptophan;
    } else if (strcmp(codon, "UAA") == 0 || strcmp(codon, "UAG") == 0 || strcmp(codon, "UGA") == 0) {
        return 7; // Stop codon
    }
    // Return an invalid amino acid for stop codons or unrecognized codons
    return 8; // Invalid amino acid
}

protein_t protein(const char *const rna) {
    size_t len = strlen(rna);
    protein_t result = { .valid = true, .count = 0 };
    for (size_t i = 0; i < len ; i += 3) {
        char codon[4]; // 3 characters + null terminator
        strncpy(codon, rna + i, 3);
        codon[3] = '\0'; // Null-terminate the codon string
        amino_acid_t amino_acid = codon_to_amino_acid(codon);
        if (amino_acid == 8) {
            result.valid = false;
            break;
        }
        if (amino_acid == 7) { // Stop codon
            break;
        }
        result.amino_acids[result.count++] = amino_acid;

    }
    return result;
}
