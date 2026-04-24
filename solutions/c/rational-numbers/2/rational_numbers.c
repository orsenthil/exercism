#include "rational_numbers.h"
#include <stdlib.h>
#include <stdint.h>
#include <math.h>

int gcd(int a, int b) {
    if (b == 0) {
        return a;
    }
    return gcd(b, a % b);
}

int power(int base, int exponent) {
    int result = 1;
    for (int i = 0; i < exponent; i++) {
        result *= base;
    }
    return result;
}

rational_t reduce(rational_t r) {
    int gcd_value = gcd(r.numerator, r.denominator);
    rational_t result;
    result.numerator = r.numerator / gcd_value;
    result.denominator = r.denominator / gcd_value;
    if (result.denominator < 0 && result.numerator > 0) {
        result.numerator = -result.numerator;
        result.denominator = -result.denominator;
    }    return result;
    return result;
}

rational_t add(rational_t a, rational_t b) {
    rational_t result;
    result.numerator = a.numerator * b.denominator + b.numerator * a.denominator;
    result.denominator = a.denominator * b.denominator;
    result = reduce(result);
    return result;
}

rational_t subtract(rational_t a, rational_t b) {
    rational_t result;
    result.numerator = a.numerator * b.denominator - b.numerator * a.denominator;
    result.denominator = a.denominator * b.denominator;
    result = reduce(result);
    return result;
}

rational_t multiply(rational_t a, rational_t b) {
    rational_t result;
    result.numerator = a.numerator * b.numerator;
    result.denominator = a.denominator * b.denominator;
    result = reduce(result);
    return result;
}

rational_t divide(rational_t a, rational_t b) {
    rational_t result;
    result.numerator = a.numerator * b.denominator;
    result.denominator = a.denominator * b.numerator;
    result = reduce(result);
    return result;
}

rational_t absolute(rational_t a) {
    rational_t result;
    result.numerator = abs(a.numerator);
    result.denominator = abs(a.denominator);
    result = reduce(result);
    return result;
}

rational_t exp_rational(rational_t a, int exponent) {
    rational_t result;
    if (exponent < 0) {
        int tmp = a.numerator;
        a.numerator = a.denominator;
        a.denominator = tmp;
        exponent = -exponent;
    }
    result.numerator = power(a.numerator, exponent);
    result.denominator = power(a.denominator, exponent);
    result = reduce(result);
    return result;
}

float exp_real(uint16_t base, rational_t exponent) {
    return powf((float) base, (float) exponent.numerator / (float) exponent.denominator);
}
