#include "all_your_base.h"

size_t rebase(int8_t digits[], int16_t input_base, int16_t output_base, size_t input_length) {
    if (input_base < 2 || output_base < 2) {
        return 0; // Invalid base
    }
    if (digits == NULL || input_length == 0) {
        return 0; // Invalid input
    }
    for (size_t i = 0; i < input_length; i++) {
        if (digits[i] < 0 || digits[i] >= input_base) {
            return 0; // Invalid digit
        }
    }
    // Convert input digits to decimal
    int32_t value = 0;
    for (size_t i = 0; i < input_length; i++) {
        value = value * input_base + digits[i];
    }
    if (value == 0) {
        digits[0] = 0;
        return 1; // Output is zero
    }
    // Convert decimal value to output base
    int8_t output_digits[DIGITS_ARRAY_SIZE];
    size_t output_length = 0;
    while (value > 0) {
        if (output_length >= DIGITS_ARRAY_SIZE) {
            return 0; // Output exceeds maximum length
        }
        output_digits[output_length++] = value % output_base;
        value /= output_base;
    }
    // Reverse the output digits
    for (size_t i = 0; i < output_length; i++) {
        digits[i] = output_digits[output_length - 1 - i];
    }
    return output_length;

}
