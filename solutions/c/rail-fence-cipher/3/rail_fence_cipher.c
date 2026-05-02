#include "rail_fence_cipher.h"
#include <string.h>
#include <stdlib.h>

static char *strcopy(const char *src) {
    char *dest = malloc(strlen(src) + 1);
    strcpy(dest, src);
    return dest;
}

char *encode(char *text, size_t rails) {
    if (rails == 1) {
        return strcopy(text);
    }
    char *encoded = malloc(strlen(text) + 1);
    size_t encoded_index = 0;
    size_t cycle = 2 * rails - 2;
    size_t len = strlen(text);
    for (size_t r = 0; r < rails; r++) {
        for (size_t i = 0; i < len; i++) {
            size_t pos = i % cycle;
            size_t rail = (pos < rails) ? pos : cycle - pos;
            if (rail == r) {
                encoded[encoded_index++] = text[i];
            }
        }
    }
    encoded[encoded_index] = '\0';
    return encoded;
}

char *decode(char *ciphertext, size_t rails) {
    if (rails == 1) {
        return strcopy(ciphertext);
    }
    // Step 1. Count the rail_len[r] for each rail and then build rail_start[rails] prefix-sum array.
    size_t rail_len[rails];
    for (size_t i = 0; i < rails; i++) {
        rail_len[i] = 0;
    }
    size_t ciphertext_len = strlen(ciphertext);
    size_t cycle = 2 * rails - 2;
    for (size_t i = 0; i < ciphertext_len; i++) {
        size_t pos = i % cycle;
        size_t rail = (pos < rails) ? pos : cycle - pos;
        rail_len[rail]++;
    }
    size_t rail_start[rails];
    rail_start[0] = 0;
    for (size_t i = 1; i < rails; i++) {
        rail_start[i] = rail_start[i - 1] + rail_len[i - 1];
    }

    // Step 2. Build the decoded string by iterating over the ciphertext and placing each character in the correct rail.
    char *decoded = malloc(ciphertext_len + 1);
    for (size_t i = 0; i < ciphertext_len; i++) {
        size_t pos = i % cycle;
        size_t rail = (pos < rails) ? pos : cycle - pos;
        /* Grab the next unused character from the rails' chunk */
        decoded[i] = ciphertext[rail_start[rail]];
        rail_start[rail]++;
    }
    decoded[ciphertext_len] = '\0';
    return decoded;
}
