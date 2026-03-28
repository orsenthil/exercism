#include "crypto_square.h"

#include <string.h>
#include <stdlib.h>
#include <ctype.h>
#include <math.h>

static void normalizeinput(const char *input, char *normalizedInput, size_t length)
{
    size_t normalizedIndex = 0;
    for (size_t index = 0; index < length; index++)
    {
        char inputChar = input[index];

        if (isalpha(inputChar) || isdigit(inputChar))
        {
            normalizedInput[normalizedIndex++] = tolower(inputChar);
        }
    }
    normalizedInput[normalizedIndex] = 0;
}

char *ciphertext(const char *input)
{
    size_t inputLength = strlen(input);
    size_t columns;
    size_t rows;

    char* normalizedInput = calloc(1, inputLength + 1);
    char* result = calloc(1, (inputLength * 2) + 1);

    normalizeinput(input, normalizedInput, inputLength);
    inputLength = strlen(normalizedInput);

    columns = (int) round(sqrt(inputLength));
    rows = columns;

    if ((rows * columns) < inputLength)
    {
        columns += 1;
    }

    memset(result, ' ', (rows * (columns + 1)));

    for(size_t colIndex = 0, normIndex = 0; (colIndex < rows) && (normIndex < inputLength); colIndex++)
    {
        for(size_t rowIndex = 0; rowIndex < columns; rowIndex++)
        {
            if(normIndex >= inputLength)
            {
                break;
            }
            result[(rowIndex * (rows + 1)) + colIndex] = normalizedInput[normIndex++];
        }
    }
    result[(columns * (rows + 1)) - 1] = 0;
    free(normalizedInput);

    return result;
}
