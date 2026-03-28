#include "perfect_numbers.h"

#include <stdint.h>

int get_factors_sum(int number) {
    int sum = 0;
    for (int i = 1; i < number; i++) {
        if (number % i == 0) {
            sum += i;
        }
    }
    return sum;
}

kind classify_number(int number) {
    int factors_sum = get_factors_sum(number);
    if (number < 1) {
        return ERROR;
    } else if (factors_sum == number) {
        return PERFECT_NUMBER;
    } else if (factors_sum < number) {
        return DEFICIENT_NUMBER;
    } else if (factors_sum > number) {
        return ABUNDANT_NUMBER;
    } else {
        return ERROR;
    }
}
