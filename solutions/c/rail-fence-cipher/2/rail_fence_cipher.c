#include "rail_fence_cipher.h"
#include <stdlib.h>
#include <string.h>

char *encode(char *text, size_t rails)
{
    int length = strlen(text);
    char* cipher = (char *) malloc(length + 1);
    int cipher_idx = 0;
    int jump1, jump2;

    // Iterate through each rail

    for(unsigned int rail = 1; rail <= rails; rail++)
    {
        // setup size of jumps (1 is downward 2 is upward.
        if (rail == 1 || rail == rails)
            jump1 = jump2 = 2 * (rails - 1);
        else {
            jump1 = 2 * (rails - rail);
            jump2 = 2 * (rail - 1);
        }
        int pos = rail - 1;
        int iteration = 0;

        // extract all members of the rail by jumping along the input

        while (pos < length)
        {
            cipher[cipher_idx++] = text[pos];
            if (iteration++ % 2 == 0)
                pos += jump1;
            else
                pos += jump2;
        }
    }

    cipher[length] = '\0';
    return cipher;
}

char *decode(char *ciphertext, size_t rails)
{
    int length = strlen(ciphertext);
    char* msg = (char *) malloc(length + 1);
    int jump1, jump2;
    int pos = 0;

    for (unsigned int rail = 1; rail <= rails ; rail++) {
        if (rail == 1 || rail == rails)
            jump1 = jump2 = 2 * (rails - 1);
        else {
            jump1 = 2 * (rails - rail);
            jump2 = 2 * (rail - 1);
        }
        int msg_idx = rail - 1;
        int iteration = 0;

        while (msg_idx < length)
        {
            msg[msg_idx] = ciphertext[pos++];
            if (iteration++ % 2 == 0)
                msg_idx += jump1;
            else
                msg_idx += jump2;
        }
    }
    msg[length] = '\0';
    return msg;
}
