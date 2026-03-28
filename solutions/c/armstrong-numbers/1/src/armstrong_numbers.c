#include "armstrong_numbers.h"
#include "math.h"

int total_digits(int num) {
    int result = 0;
    while (num > 0) {
        result += 1;
        num = num / 10;
    }
    return result;
}

bool is_armstrong_number(int num) {
    int orig = num;
    int n = 0;
    int sum = 0;
    int d = 0;

    n = total_digits(num);

    while (num > 0) {
        d = num % 10;
        sum += pow(d, n);
        num = num / 10;
    }

    return sum == orig;
}