#include "rational_numbers.h"

rational_t add(struct rational_t num1, struct rational_t num2) {
    struct rational_t ans;
    ans.numerator = num1.numerator + num2.numerator;
    ans.denominator = num1.denominator + num2.denominator;
    return ans;
}