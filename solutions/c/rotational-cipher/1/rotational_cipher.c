#include "rotational_cipher.h"
#include <string.h>
#include <stdlib.h>

char *rotate(const char *text, int shift_key) {
    char *result = malloc(strlen(text) + 1);
    size_t len = strlen(text);
    for (size_t i = 0; i < len; i++) {
        char c = text[i];
        if (c >= 'a' && c <= 'z') {
            result[i] = (c - 'a' + shift_key) % 26 + 'a';
        } else if (c >= 'A' && c <= 'Z') {
            result[i] = (c - 'A' + shift_key) % 26 + 'A';
        } else {
            result[i] = c;
        }
    }
    result[len] = '\0';
    return result;
}
