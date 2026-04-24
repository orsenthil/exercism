#include "phone_number.h"
#define _POSIX_C_SOURCE 200809L 
#include <stdlib.h>
#include <string.h>

char *phone_number_clean(const char *input) {
    char *output = malloc(strlen(input) + 1);
    size_t output_index = 0;
    for (size_t i = 0; i < strlen(input); i++) {
        if (input[i] >= '0' && input[i] <= '9') {
            output[output_index++] = input[i];
        }
    }
    if (output_index == 11 && output[0] == '1') {
        memmove(output, output + 1, output_index - 1);
        output_index--;
    }
    if (output_index != 10) {
        return strdup("0000000000");
    }
    if (output[0] == '0' || output[0] == '1') {
        return strdup("0000000000");
    }
    if (output[3] == '0' || output[3] == '1') {
        return strdup("0000000000");
    }
    output[output_index] = '\0';
    return output;
}
