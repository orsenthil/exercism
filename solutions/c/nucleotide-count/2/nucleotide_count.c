#include "nucleotide_count.h"
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

char *count(const char *dna_strand) {
    int count_A = 0;
    int count_C = 0;
    int count_G = 0;
    int count_T = 0;

    size_t buf_size = 20; // "A:0 C:0 G:0 T:0" is 19 characters + null terminator
    char *buf = malloc(buf_size);

    if (buf == NULL) {
        return NULL;
    }

    for (size_t i = 0; i < strlen(dna_strand); i++) {
        if (dna_strand[i] != 'A' && dna_strand[i] != 'C' && dna_strand[i] != 'G' && dna_strand[i] != 'T') {
            free(buf);
            char *empty = malloc(1);
            empty[0] = '\0';
            return empty;
        }
        switch (dna_strand[i]) {
            case 'A':
                count_A++;
                break;
            case 'C':
                count_C++;
                break;
            case 'G':
                count_G++;
                break;
            case 'T':
                count_T++;
                break;
        }
    }
    snprintf(buf, buf_size, "A:%d C:%d G:%d T:%d", count_A, count_C, count_G, count_T);
    return buf;
}
