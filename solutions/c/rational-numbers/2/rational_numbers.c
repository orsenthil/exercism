#include "rational_numbers.h"
#include <math.h>
#include <stdlib.h>

int _gcd_euclidean(int a, int b);

rational_t add(rational_t r1, rational_t r2) {
    rational_t  ret;

    ret.numerator = r1.numerator * r2.denominator + r2.numerator * r1.denominator;
    ret.denominator = r1.denominator * r2.denominator;

    return reduce(ret);
}

rational_t subtract(rational_t r1, rational_t r2) {
    rational_t ret;

    ret.numerator = r1.numerator * r2.denominator - r2.numerator * r1.denominator;
    ret.denominator = r1.denominator * r2.denominator;

    return reduce(ret);
}

rational_t multiply(rational_t r1, rational_t r2) {
    rational_t ret;

    ret.numerator = r1.numerator * r2.numerator;
    ret.denominator = r1.denominator * r2.denominator;

    return reduce(ret);
}

rational_t divide(rational_t r1, rational_t r2) {
    rational_t ret;

    ret.numerator = r1.numerator * r2.denominator;
    ret.denominator = r2.numerator * r1.denominator;

    return reduce(ret);
}

rational_t absolute(rational_t r) {
    rational_t ret = {abs(r.numerator), abs(r.denominator)};
    return reduce(ret);
}

rational_t exp_rational(rational_t r, uint16_t n) {
    rational_t ret = {pow(r.numerator, n), pow(r.denominator, n)};
    return reduce(ret);
}

float exp_real(float x, rational_t r) {
    return pow(pow(x, r.numerator), 1.0 / r.denominator);
}

int _gcd_euclidean(int a, int b) {
    int larger = a >= b ? a : b;
    int smaller = a <= b ? a : b;

    return smaller == 0 ? larger : _gcd_euclidean(smaller, larger % smaller );
}

rational_t reduce(rational_t r) {
    int gcd = _gcd_euclidean(abs(r.numerator), abs(r.denominator));

    rational_t ret = {r.numerator / gcd, r.denominator / gcd};

    if (ret.denominator < 0) {
        ret.numerator *= -1;
        ret.denominator *= -1;
    }

    return ret;
}