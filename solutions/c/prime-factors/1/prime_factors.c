#include "prime_factors.h"

size_t find_factors(uint64_t n, uint64_t factors[static MAXFACTORS]) {
    if (n == 1) {
        return 0;
    }
    size_t count = 0;
    uint64_t factor = 2;
    while (factor * factor <= n) {
        while (n % factor == 0) {
            if (count < MAXFACTORS) {
                factors[count++] = factor;
            }
            n /= factor;
        }
        factor++;   
    }
    if (n > 1 && count < MAXFACTORS) {
        factors[count++] = n;
    }
    return count;
}
