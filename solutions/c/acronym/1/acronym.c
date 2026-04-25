#include "acronym.h"
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <stdbool.h>

char *abbreviate(const char *phrase) {
    if (phrase == NULL) {
        return NULL;
    }
    if (strlen(phrase) == 0) {
        return NULL;
    }
    size_t len = strlen(phrase);
    char *abbreviation = malloc(len + 1);
    bool in_word = false;
    size_t output_index = 0;
    for (size_t i = 0; i < len; i++) {
        char c = phrase[i];
        if (isalpha(c) && !in_word) {
            abbreviation[output_index] = toupper(c);
            output_index++;
            in_word = true;
        } else if (c == ' ' || c == '-' || c == '_') {
            in_word = false;
        }
    }
    abbreviation[output_index] = '\0';
    return abbreviation;
}
