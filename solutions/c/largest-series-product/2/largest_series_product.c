#include "largest_series_product.h"

#include <ctype.h>
#include <string.h>

int64_t largest_series_product(char *digits, size_t span) {
    size_t digits_length = strlen(digits);

    if (span > digits_length)
        return -1;

    int64_t result = 0;

    // If there is no span, there is no possiblity for all candidate being 0,
    // so result is 1

    if (span == 0)
        result = 1;

    for(size_t i = 0; i <= digits_length - span; i++) {
        int64_t candidate = 1;

        for(size_t j = i; j < i + span; j++) {
            if (!isdigit(digits[j]))
                return  -1;
            candidate *= (digits[j] - '0');
        }

        if (candidate > result)
            result = candidate;
    }
    return result;
}
