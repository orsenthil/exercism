#include "armstrong_numbers.h"
#include <math.h>

int get_digits(int candidate) {
    int digits = 0;

    do {
        digits += 1;
        candidate /= 10;
    } while (candidate > 0);

    return digits;
}

bool is_armstrong_number(int candidate) {
    int original = candidate;
    int digits = get_digits(candidate);
    int digit = 0;
    int sum = 0;

    while (candidate > 0) {
        digit = candidate % 10;
        sum += pow(digit, digits);
        candidate /= 10;
    }
    return sum == original;    
}