#include "acronym.h"

#include <stdlib.h>


char *abbreviate(const char *phrase)
{
    if (phrase == NULL || strcmp(phrase, "") == 0) {
        return NULL;
    }

    return "";
}