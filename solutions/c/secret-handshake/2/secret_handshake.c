#include "secret_handshake.h"

#include <stdlib.h>

#define MAXEVENTS 4

static char *events[] = {
        "wink", "double blink", "close your eyes", "jump"
};

const char **commands(size_t number)
{
    const char **handshake = NULL;
    int count = 0;

    handshake = realloc(handshake, count * sizeof(char *));
    handshake[0] = NULL;

    for (int i = 0; i < MAXEVENTS; i++)
    {
        if (number & (0x1 << i))
        {
            count += 1;
            handshake = realloc(handshake, count * sizeof(char *));
            handshake[count - 1] = events[i];
        }
    }

    if (count != 0 && (number & (0x1 << MAXEVENTS)))
        for (int i = 0; i < count / 2; i++)
        {
            const char *tmp = handshake[i];
            handshake[i] = handshake[count - (i + 1)];
            handshake[count - (i + 1)] = tmp;
        }

    return handshake;
}
