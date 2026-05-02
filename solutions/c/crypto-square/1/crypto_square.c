#include "crypto_square.h"
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <math.h>

char *ciphertext(const char *input) {
    char *result = malloc(strlen(input) + 1);
    // Step 1. Normalize the input
    size_t input_index = 0;
    size_t result_index = 0;
    while (input[input_index] != '\0') {
        if (isalnum(input[input_index])) {
            result[result_index] = tolower(input[input_index]);
            result_index++;
        }
        input_index++;
    }
    result[result_index] = '\0';

    size_t len = strlen(result);

    if (len == 0) {
        return result;
    }

    // Step 2. Determine the dimensions of the rectangle
    size_t c = ceil(sqrt(len));
    size_t r = ceil((double)len / c);

    // Step 3. Fill the rectangle with the normalized characters
    char rectangle[r][c];
    size_t index = 0;
    for (size_t i = 0; i < r; i++) {
        for (size_t j = 0; j < c; j++) {
            if (index < len) {
                rectangle[i][j] = result[index];
                index++;
            } else {
                rectangle[i][j] = ' ';
            }
        }
    }

    char *output = malloc((c * r + c - 1) + 1);
    for (size_t j = 0; j < c; j++) {
        for (size_t i = 0; i < r; i++) {
            output[j * (r + 1) + i] = rectangle[i][j];
        }
        output[j * (r + 1) + r] = ' ';
    }
    output[(c * (r + 1)) - 1] = '\0';
    free(result);
    return output;

}
