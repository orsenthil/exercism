#include "nth_prime.h"
#include <stdbool.h>

static bool is_prime(uint32_t num) {
    if (num < 2) {
        return false;
    }
    if (num == 2) {
        return true;
    }
    if (num % 2 == 0) {
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
    if (n == 1) {
        return 2;
    }
    uint32_t count = 1;
    uint32_t candidate = 3;

    while (count < n) {
        if (is_prime(candidate)) {
            count++;
            if (count == n) {
                return candidate;
            }
        }
        if (count < n)
        {
            candidate += 2; // Skip even numbers
        }
    }

    return 0; // This line will never be reached, but it prevents a compiler warning.
}
