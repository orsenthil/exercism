#include "etl.h"
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

int cmp(const void *a, const void *b) {
    return ((new_map *)a)->key - ((new_map *)b)->key;
}

int convert(const legacy_map *input, const size_t input_len, new_map **output) {
    size_t total_len = 0;
    for (size_t i = 0; i < input_len; i++) {
        total_len += strlen(input[i].keys);
    }
    *output = calloc(total_len, sizeof(new_map));
    size_t output_len = 0;
    for (size_t i = 0; i < input_len; i++) {
        const char *keys = input[i].keys;
        size_t keys_len = strlen(keys);
        for (size_t j = 0; j < keys_len; j++) {
            (*output)[output_len].key = tolower(keys[j]);
            (*output)[output_len].value = input[i].value;
            output_len++;
        }
    }
    qsort(*output, output_len, sizeof(new_map), cmp);
    return output_len;
}
