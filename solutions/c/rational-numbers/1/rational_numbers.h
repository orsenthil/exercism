#ifndef RATIONAL_NUMBERS_H
#define RATIONAL_NUMBERS_H

struct rational_t {
    int numerator;
    int denominator;
};

rational_t add(rational_t num1, rational_t num2);

#endif
