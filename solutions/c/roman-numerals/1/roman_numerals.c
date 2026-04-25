#include "roman_numerals.h"
#include <stdlib.h>
#include <string.h>

char *to_roman_numeral(unsigned int number) {
    unsigned int values [] = {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1};
    char *numerals[] = {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"};
    char *result = malloc(16);

    for (size_t i = 0; i < 13; i++) {
        while (number >= values[i]) {
            strcat(result, numerals[i]);
            number -= values[i];
        }
    }
    return result;
}
