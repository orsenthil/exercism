#include "hamming.h"
#include <string.h>


int compute(const char *lhs, const char *rhs) {
    if (strlen(lhs) != strlen(rhs)) {
        return -1;
    }
    int count = 0;

    for (size_t i = 0; i < strlen(lhs); i++) {
        if (lhs[i] != rhs[i]) {
            count += 1;
        }
    }
    return count;
}

