#include "scrabble_score.h"

#include <ctype.h>

int SCORES[] = {1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3,
                1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10};

unsigned int score(const char *word) {
    int score = 0;

    while (*word) {
        if (isalpha(*word)) {
            score += SCORES[tolower(*word) - 'a'];
        }

        word++;
    }
    return score;
}
