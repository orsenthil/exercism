#include "largest_series_product.h"
#include <string.h>

int64_t largest_series_product(char *digits, size_t span) {
    if (span > strlen(digits)) {
        return -1;
    }

    int64_t max_product = 0;
    for (size_t i = 0; i <= strlen(digits) - span; i++) {
        int64_t product = 1;
        for (size_t j = 0; j < span; j++) {
            if (digits[i + j] < '0' || digits[i + j] > '9') {
                return -1;
            }
            product *= digits[i + j] - '0';
        }
        if (product > max_product) {
            max_product = product;
        }
    }
    return max_product;
}
