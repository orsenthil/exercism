#include "nth_prime.h"
#include <stdbool.h>

static bool is_prime(uint32_t num) {
    if (num < 2) {
        return false;
    }
    for (uint32_t i = 2; i * i <= num; i++) {
        if (num % i == 0) {
            return false;
        }
    }
    return true;
}

uint32_t nth(uint32_t n) {
    if (n == 0) {
        return 0;
    }
    uint32_t count = 0;
    uint32_t candidate = 2;

    while (count < n) {
        if (is_prime(candidate)) {
            count++;
            if (count == n) {
                return candidate;
            }
        }
        candidate++;
    }

    return 0; // This line will never be reached, but it prevents a compiler warning.
}
