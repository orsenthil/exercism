#include "isogram.h"
#include <string.h>
#include <stdint.h>
#include <ctype.h>

bool is_isogram(const char phrase[]) {
    if (phrase == NULL) {
        return false;
    }

    uint8_t length = strlen(phrase);

    for (uint8_t i = 0; i < length - 1; i++) {
        for (uint8_t j = i + 1; j < length; j++) {
            if (tolower(phrase[i]) == tolower(phrase[j]) && isalpha(phrase[i]) && isalpha(phrase[j])) {
                return false;
            }
        }
    }
    return true;
}