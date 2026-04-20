#include "luhn.h"
#include <stdlib.h>
#include <ctype.h>
#include <string.h>

bool luhn(const char *num) {
    if (num == NULL) {
        return false;
    }
    char *copy = malloc(strlen(num) + 1);
    size_t j = 0;
    for (size_t i = 0; i < strlen(num); i++) {
     if (num[i] == ' ') {
        continue;
     }
     if (isdigit((int)num[i])) {
        copy[j] = num[i];
        j++;
     } else {
        free(copy);
        return false;
     }
    }
    copy[j] = '\0';
    size_t len = strlen(copy);
    if (len <= 1) {
        free(copy);
        return false;
    }
    bool double_digit = false;
    int result = 0;
    for (int i = (int)len - 1; i >= 0; i--) {
        if (double_digit) {
            int digit = copy[i] - '0';
            digit = digit * 2;
            if (digit > 9) {
                digit = digit - 9;
            }
            result += digit;
            double_digit = false;
        } else {
            result += copy[i] - '0';
            double_digit = true;
        }


    }

    return result % 10 == 0;
}
