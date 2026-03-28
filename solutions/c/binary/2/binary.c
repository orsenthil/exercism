#include "binary.h"
#include <ctype.h>

const int base = 2;

int convert(const char *input) {
    int res = 0;

    while (*input != '\0')
    {
        if (!isdigit(*input))
            return INVALID;

        int n  = *input - '0';
        if (n > 1) {
            return INVALID;
        }

        res = res * base + n;
        input += 1;

    }
    return res;
}
