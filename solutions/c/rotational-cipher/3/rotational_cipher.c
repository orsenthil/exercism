#include "rotational_cipher.h"

#include <stdlib.h>
#include <string.h>
#include <ctype.h>

char *rotate(const char *text, int shift_key) {
    int len = strlen(text);
    char *result = malloc(strlen(text) + 1);
    for (int i = 0; i < len; ++i) {
        if (isalpha(text[i])) {
            char base = isupper(text[i]) ? 'A' : 'a';
            result[i] = base + (text[i] - base + shift_key) % 26;
        } else {
            result[i] = text[i];
        }
    }
    return result;
}
