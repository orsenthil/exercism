#include "all_your_base.h"

size_t rebase(int8_t* digits, int16_t input_base, int16_t output_base, int16_t input_length) {
    if (input_length == 0) return 0;
    if (input_base <= 1) return 0;
    if (output_base <= 1) return 0;

    for (int i = 0; i < input_length; i++) {
        if (digits[i] < 0)
            return 0;
        if (digits[i] >= input_base)
            return 0;
    }

    int num = digits[input_length - 1];
    int base = input_base;

    for (int i = input_length - 2; i >= 0; i--) {
        num += base * digits[i];
        base *= input_base;
    }

    int i = 0;
    int8_t rbdigits[DIGITS_ARRAY_SIZE];

    while (num >= output_base) {
        rbdigits[i++] = num % output_base;
        num = num / output_base;
    }

    rbdigits[i] = num;
    int k = 0;

    for (int j = i; j >= 0; j--) {
        digits[j] = rbdigits[k++];
    }

    return k;
}
