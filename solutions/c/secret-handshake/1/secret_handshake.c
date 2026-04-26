#include "secret_handshake.h"
#include <stdlib.h>

const char **commands(size_t number) {
    const char *commands[4] = { "wink", "double blink", "close your eyes", "jump" };
    const char **result = calloc(4, sizeof(char *));
    size_t count = 0;
    for (size_t i = 0; i < 4; i++) {
        if (number & (1 << i)) {
            result[count++] = commands[i];
        }
    }
    if (number & (1 << 4) && count > 0) {
        size_t left = 0, right = count - 1;
        while (left < right) {
            const char *temp = result[left];
            result[left] = result[right];
            result[right] = temp;
            left++;
            right--;
        }
    }
    return result;
}
