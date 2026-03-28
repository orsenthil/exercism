#include "armstrong_numbers.h"
#include <math.h>

int get_digits(int number) {
    int digits = 0;
    do {
        digits += 1;
        number /= 10;
    } while (number > 0);

    return digits;
}

bool is_armstrong_number(int number) {
    int original = number;
    int digits = get_digits(number);

    int digit = 0;
    int sum = 0;

    while (number > 0) {
        digit = number % 10;
        sum += pow(digit, digits);
        number /= 10;
    }

    return sum == original;
}