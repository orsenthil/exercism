#include "bob.h"

#include <ctype.h>

char *hey_bob(char *greeting)
{
    int upper = 0;
    int lower = 0;
    int ch = 0;

    char* p = greeting;

    while (*p != '\0') {
        if (isupper(*p))
            upper++;
        else if (islower(*p))
            lower++;
        else if (!isspace(*p))
            ch++;

        p++;
    }

    do {
        p--;
    } while(isspace(*p));

    if (*p == '?') {
        if (upper > 0 && lower == 0)
            return "Calm down, I know what I'm doing!";
        else
            return "Sure.";
    }

    if (lower == 0) {
        if (upper > 0)
            return "Whoa, chill out!";
        else if (ch == 0)
            return "Fine. Be that way!";
    }

    return "Whatever.";

}
