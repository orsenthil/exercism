#ifndef ROMAN_NUMERALS_H
#define ROMAN_NUMERALS_H

typedef struct RomanNumeral
{
    unsigned int numeral;
    char* roman;
    unsigned int len;
} RomanNumeral;

char *to_roman_numeral(unsigned int number);

#endif
