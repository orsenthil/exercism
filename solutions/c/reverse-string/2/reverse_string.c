#include "reverse_string.h"
#include <string.h>
#include <stdlib.h>

char *reverse(const char *value) {
    char *reversed = malloc(strlen(value) + 1);
    int i, j;
    for (i = strlen(value) - 1, j = 0; i >= 0; i--, j++) {
        reversed[j] = value[i];
    }
    reversed[j] = '\0';
    return reversed;
}
