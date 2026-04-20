#include "sieve.h"
#include <stdbool.h>
#include <string.h>

/// Calculate at most `max_primes` prime numbers in the interval [2,limit]
/// using the Sieve of Eratosthenes and store the prime numbers in `primes`
/// in increasing order.
/// The function returns the number of calculated primes.
uint32_t sieve(uint32_t limit, uint32_t *primes, size_t max_primes) {
    if (limit < 2) {
        return 0;
    }

    bool is_composite[limit + 1];
    memset(is_composite, 0, sizeof(is_composite));

    uint32_t count = 0;
    for (uint32_t p = 2; p <= limit; p++) {
        if (!is_composite[p]) {
            primes[count++] = p;
            if (count >= max_primes) {
                break;
            }
            for (uint32_t j = p * p; j <= limit; j += p) {
                is_composite[j] = true;
            }
        }
    }   
    return count;
}
