#include "zebra_puzzle.h"

#include <stdlib.h>
#include <stddef.h>
#include <assert.h>
#include <stdbool.h>



enum {
    PERMUTATION_SIZE = 5,
    PERMUTATION_COUNT = 120
};

enum {
    ENGLISH = 0,
    JAPANESE,
    NORWEGIAN,
    SPANISH,
    UKRAINIAN,

    BLUE = 0,
    GREEN,
    IVORY,
    RED,
    YELLOW,

    DOG = 0,
    FOX,
    HORSE,
    SNAILS,
    ZEBRA,

    COFFEE = 0,
    MILK,
    ORANGE_JUICE,
    TEA,
    WATER,

    CHESTERFIELD = 0,
    KOOL,
    LUCKY_STRIKE,
    OLD_GOLD,
    PARLIAMENT,
};

static const char* nationality_string[] = {
    "English",
    "Japanese",
    "Norwegian",
    "Spanish",
    "Ukrainian"
};

typedef struct {
    int _[PERMUTATION_COUNT][PERMUTATION_SIZE];
} permutations_t;

static permutations_t all_permutations(void) {
    permutations_t all;
    int count = 0;
    for (int a = 0; a < PERMUTATION_SIZE; a++) {
        for (int b = 0; b < PERMUTATION_SIZE; b++) {
            if (b == a) continue;
            for (int c = 0; c < PERMUTATION_SIZE; c++) {
                if (c == a || c == b) continue;
                for (int d = 0; d < PERMUTATION_SIZE; d++) {
                    if (d == a || d == b || d == c) continue;
                    for (int e = 0; e < PERMUTATION_SIZE; e++) {
                        if (e == a || e == b || e == c || e == d) continue;
                        all._[count][0] = a;
                        all._[count][1] = b;
                        all._[count][2] = c;
                        all._[count][3] = d;
                        all._[count][4] = e;
                        count++;
                    }
                }
            }
        }
    }
    return all;
}

solution_t solve_puzzle(void) {
    permutations_t permutations = all_permutations();
    for (size_t c = 0; c < PERMUTATION_COUNT; c++) {
        int* color = permutations._[c];
        if (color[GREEN] - 1 != color[IVORY]) continue;
        for (size_t n = 0; n < PERMUTATION_COUNT; n++) {
            int* nationality = permutations._[n];

            // The Norwegian lives in the first house.
            if (nationality[NORWEGIAN] != 0) continue;
            // The Englishman lives in the red house.
            if (nationality[ENGLISH] != color[RED]) continue;
            // The Norwegian lives next to the blue house.
            if (abs(nationality[NORWEGIAN] - color[BLUE]) != 1) continue;

            for (size_t p = 0; p < PERMUTATION_COUNT; p++) {
                int* pet = permutations._[p];
                // The Spaniard owns the dog.
                if (nationality[SPANISH] != pet[DOG]) continue;
                for (size_t beverage_i = 0; beverage_i < PERMUTATION_COUNT; beverage_i++) {
                    int* beverage = permutations._[beverage_i];
                    // Coffee is drunk in the green house.
                    if (beverage[COFFEE] != color[GREEN]) continue;
                    // The Ukrainian drinks tea.
                    if (nationality[UKRAINIAN] != beverage[TEA]) continue;
                    // Milk is drunk in the middle house.
                    if (beverage[MILK] != 2) continue;

                    for (size_t b = 0; b < PERMUTATION_COUNT; b++) {
                        int* brand = permutations._[b];

                        // The Old Gold smoker owns snails.
                        if (brand[OLD_GOLD] != pet[SNAILS]) continue;
                        // Kools are smoked in the yellow house.
                        if (brand[KOOL] != color[YELLOW]) continue;

                        // The man who smokes Chesterfields lives in the house next to the man with the fox.
                        if (abs(brand[CHESTERFIELD] - pet[FOX]) != 1) continue;
                        // Kools are smoked in the house next to the house where the horse is kept.
                        if (abs(brand[KOOL] - pet[HORSE]) != 1) continue;
                        // The Lucky Strike smoker drinks orange juice.
                        if (brand[LUCKY_STRIKE] != beverage[ORANGE_JUICE]) continue;
                        // The Japanese smokes Parliaments.
                        if (nationality[JAPANESE] != brand[PARLIAMENT]) continue;

                        return (solution_t) {
                            .drinks_water = nationality_string[nationality[beverage[WATER]]],
                            .owns_zebra = nationality_string[nationality[pet[ZEBRA]]]
                        };
                    }
                }
            }
        }
    }
    assert(!"No solution found");
}
