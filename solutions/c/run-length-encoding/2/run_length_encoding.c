#include "run_length_encoding.h"

#include <ctype.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

char *encode(const char *text)
{
    char* result = calloc(strlen(text) + 1, sizeof(char));
    int count = 0;
    char last = '\0';

    for(size_t i = 0; i < strlen(text); i++) {
        if (last == '\0')
        {
            last = text[i];
            count = 1;
        } else if (text[i] == last)
        {
            count += 1;
        } else if (text[i] != last)
        {
            char buffer[4];

            if (count > 1)
                sprintf(buffer, "%d%c", count, last);
            else
                sprintf(buffer, "%c", last);

            strcat(result, buffer);

            // Update last
            last = text[i];
            count = 1;
        }
    }

    char buffer[4];
    if (count > 1)
        sprintf(buffer, "%d%c", count, last);
    else
        sprintf(buffer, "%c", last);

    strcat(result, buffer);

    return result;
}

const int RESULT_CAPACITY = 1000;

char* decode(const char *data) {
    char* result = calloc(RESULT_CAPACITY, sizeof(char));
    size_t result_length = 0;
    for (size_t i = 0; i < strlen(data); i++)
    {
        int counter = 1;
        if (isdigit(data[i])) {
            sscanf(&data[i], "%d", &counter);
            while (isdigit(data[i])) i++;
        }

        for(int j = 0; j < counter; j++)
            result[result_length++] = data[i];
    }
    return result;
}
