#include "diamond.h"

#include <stdlib.h>
#include <string.h>
#include <stdio.h>

#define LETTER_LENGTH 26

static char letters[LETTER_LENGTH] = {
        'A', 'B', 'C', 'D', 'E', 'F',
        'G', 'H', 'I', 'J', 'K', 'L',
        'M', 'N', 'O', 'P', 'Q', 'R',
        'S', 'T', 'U', 'V', 'W', 'X',
        'Y', 'Z'
};

static void populate_string(char** diamond, int i, char c, int length, int side, int middle);

char **make_diamond(const char letter)
{

    if (letter < 'A' || letter > 'Z') {
        return NULL;
    }

    int letter_index = 0, i;

    while (letter != letters[letter_index++]);

    int length = (2 * letter_index) - 1, side, middle;

    char** diamond = calloc(length, sizeof(char*));

    for(i = 0; i < letter_index; i++) {
        side = letter_index - 1 -i;
        middle = (2 * i) - 1;

        populate_string(diamond, i, letters[i], length, side, middle);

        if (i != length - i - 1) {
            populate_string(diamond, length - i - 1, letters[i], length, side, middle);
        }
    }

    return diamond;
}

void free_diamond(char** diamond)
{
    if (diamond == NULL) {
        return;
    }
    free(diamond);
}

void populate_string(char** diamond, int i, char c, int length, int side, int middle)
{
    diamond[i] = calloc(length, sizeof(char));

    int pos = 0;

    memset(diamond[i] + pos, ' ', side);

    pos += side;

    sprintf(diamond[i] + pos, "%c", c);

    pos += 1;

    if (middle > 0)
    {
        memset(diamond[i] + pos, ' ', middle);
        pos += middle;

        sprintf(diamond[i] + pos, "%c", c);
        pos += 1;
    }

    memset(diamond[i] + pos, ' ', side);

    pos += side;
}
