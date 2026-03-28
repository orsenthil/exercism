#include "acronym.h"

#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <stdint.h>


char *abbreviate(const char *phrase)
{
    uint32_t count = 1;
    size_t j = 0;

    if (phrase == NULL || *phrase == 0)
    {
        return NULL;
    }
    else
    {
        for(long unsigned int i = 0; i < strlen(phrase); ++i)
        {
            if (phrase[i] == ' ')
            {
                ++count;
            }
        }
    }

    char *acronym = calloc(count + 1, sizeof(*acronym));

    for (long unsigned int i = 0; i < strlen(phrase) - 1; i++)
    {
        if (i == 0 || (phrase[i-1] == '-' && phrase[i] != ' ') || (phrase[i-1] == ' ' && phrase[i] != '_') || (phrase[i-1] == '_' && phrase[i] != ' ') )
        {
            if (phrase[i] != '-')
            {
                acronym[j] = toupper(phrase[i]);
                j += 1;
            }
        }
    }
    return acronym;
}
