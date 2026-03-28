#include "nth_prime.h"

uint32_t nth(uint32_t n) {
    uint32_t primes[10000];
    uint32_t cmp = 0;
    primes[cmp] = 2;
    uint32_t number = 1;

    if (n == 0) {
        return 0;
    } else if (n == 1) {
        return 2;
    } else {
        while (cmp < (n - 1))
        {
            START:
            number = number + 2;
            for (uint32_t iprime = 0; iprime <= cmp; iprime++)
            {
                if (number % primes[iprime] == 0)
                {
                    goto START;
                }
            }
            cmp++;
            primes[cmp]= number;
        }
    }
    return number;
}
