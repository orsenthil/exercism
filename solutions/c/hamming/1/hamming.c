#include "hamming.h"

#include <string.h>

int compute(const char *lhs, const char *rhs) {
    if (strlen(lhs) != strlen(rhs)) {
        return -1;
    }
    int distance = 0;
    int i = 0;
    while (lhs[i] != '\0') {
        if (lhs[i] != rhs[i]) {
            distance++;
        }
        i++;
    }
    return distance;
}
