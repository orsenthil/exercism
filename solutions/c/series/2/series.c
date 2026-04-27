#include "series.h"
#include <string.h>
#include <stdlib.h>

slices_t slices(char *input_text, unsigned int substring_length) {
    size_t len = strlen(input_text);
    if (substring_length > len || substring_length == 0) {
        return (slices_t){0, NULL};
    }
    slices_t result = {0, NULL};
    result.substring_count = len - substring_length + 1;
    result.substring = calloc(result.substring_count, sizeof(char *));
    for (size_t i = 0; i < result.substring_count; i++) {
        result.substring[i] = malloc(substring_length + 1);
        strncpy(result.substring[i], input_text + i, substring_length);
        result.substring[i][substring_length] = '\0';
    }
    return result;
}
