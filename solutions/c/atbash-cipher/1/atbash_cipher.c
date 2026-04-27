#include "atbash_cipher.h"
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

static char atbash_char(char c) {
    c = tolower(c);
    if (c >= 'a' && c <= 'z') {
        return 'z' - (c - 'a');
    }
    if (c >= '0' && c <= '9') {
        return c;
    }
    return 0;
}

char *atbash_encode(const char *input) {
    char *encoded = malloc((4 * strlen(input) + 1));
    size_t encoded_index = 0;
    size_t input_index = 0;
    size_t group_index = 0;
    while (input[input_index] != '\0') {
        char c = atbash_char(input[input_index]);
        if (c != 0) {
            if (group_index > 0 && group_index % 5 == 0) {
                encoded[encoded_index++] = ' ';
            }
            encoded[encoded_index++] = c;
            group_index++;
        }
        input_index++;

    }
    encoded[encoded_index] = '\0';
    return encoded;
}

char *atbash_decode(const char *input) {
    char *decoded = malloc((strlen(input) + 1));
    size_t decoded_index = 0;
    size_t input_index = 0;
    while (input[input_index] != '\0') {
        char c = atbash_char(input[input_index]);
        if (c != 0) {
            decoded[decoded_index++] = c;
        }
        input_index++;
    }
    decoded[decoded_index] = '\0';
    return decoded;
}
