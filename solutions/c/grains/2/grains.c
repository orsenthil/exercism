#include "grains.h"
#include <stdio.h>

uint64_t square(uint8_t index) {
    if (index == 0) {
        return 0;
    }
    uint64_t total = 1;
    for (int i = 2; i <= index; i++) {
        total *= 2;
    }
    return total;
}

uint64_t total(void) {
    uint64_t total = 1;
    uint64_t cumm = 1;

    for (int i = 2; i <= 64; i++) {
        total *= 2;
        cumm += total;
    }

    return cumm;
}