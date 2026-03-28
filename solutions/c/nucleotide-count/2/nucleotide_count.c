#include "nucleotide_count.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *count(const char *dna_strand) {
    char *ret = malloc(30);
    int numbers[4] = {0};

    for (; *dna_strand; dna_strand++)
    {
        if (*dna_strand == 'A')
            numbers[0] += 1;
        else if (*dna_strand == 'C')
            numbers[1] += 1;
        else if (*dna_strand == 'G')
            numbers[2] += 1;
        else if (*dna_strand == 'T')
            numbers[3] += 1;
        else {
            strcpy(ret, "");
            return ret;
        }
    }

    sprintf(ret, "%c:%d %c:%d %c:%d %c:%d", 'A', numbers[0], 'C', numbers[1], 'G', numbers[2], 'T', numbers[3]);

    return ret;
}
