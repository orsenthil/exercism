#include "scrabble_score.h"

#include <stdint.h>
#include <stdlib.h>


unsigned int score(const char *word) {
    static const int letter_scores[26] = {1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10};
    unsigned int total_score = 0;
    for (size_t i = 0; word[i] != '\0'; i++) {
        char c = word[i];
        if (c >= 'a' && c <= 'z') {
            total_score += letter_scores[c - 'a'];
        } else if (c >= 'A' && c <= 'Z') {
            total_score += letter_scores[c - 'A'];
        }
    }
    return total_score;
}
