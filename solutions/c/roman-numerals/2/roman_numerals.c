#include "roman_numerals.h"

#include <stdlib.h>
#include <string.h>

static const RomanNumeral lookup[] = {
        {1000, "M", 1},
        {900, "CM", 2},
        {500, "D", 1},
        {400, "CD", 2},
        {100, "C", 1},
        {90, "XC", 2},
        {50, "L", 1},
        {40, "XL", 2},
        {10, "X", 1},
        {9, "IX", 2},
        {5, "V", 1},
        {4, "IV", 2},
        {1, "I", 1}
};

char *to_roman_numeral(unsigned int number)
{
    // longest roman numeral is MMMDCCCLXXXVIII; 15 + 1
    char* output = calloc(16, 1);
    char* start_output = output;

    if (!output) {
        return NULL;
    }

    for(int i = 0; i < 13; i++)
    {
        RomanNumeral numeral= lookup[i];

        while (number >= numeral.numeral)
        {
            memcpy(output, numeral.roman, numeral.len);
            output += numeral.len;
            number -= numeral.numeral;
        }
    }
    return start_output;
}
