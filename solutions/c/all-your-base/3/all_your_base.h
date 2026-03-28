#ifndef ALL_YOUR_BASE_H
#define ALL_YOUR_BASE_H

#include <stddef.h>
#include <stdint.h>
#include <stdio.h>

#define DIGITS_ARRAY_SIZE 64

size_t rebase(int8_t* digits, int16_t input_base, int16_t output_base, int16_t input_length);

#endif
