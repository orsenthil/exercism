#include "pythagorean_triplet.h"
#include <stdlib.h>
#include <stddef.h>

triplets_t *triplets_with_sum(uint16_t sum) {
    triplets_t *triplets = malloc(sizeof(triplets_t));
    triplets->count = 0;
    triplets->triplets = NULL;
    for (uint16_t a = 1; a < sum; a++) {
        for (uint16_t b = a + 1; b < sum; b++) {
            uint16_t c = sum - a - b;
            if (c > b && a * a + b * b == c * c) {
                triplets->count++;
                triplets->triplets = realloc(triplets->triplets, triplets->count * sizeof(triplet_t));
                triplets->triplets[triplets->count - 1].a = a;
                triplets->triplets[triplets->count - 1].b = b;
                triplets->triplets[triplets->count - 1].c = c;
            }
        }
    }
    return triplets;
}                                                                     
void free_triplets(triplets_t *triplets)
{
    free(triplets->triplets);
    free(triplets);
}
