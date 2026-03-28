#include "eliuds_eggs.h"

int egg_count(int bits) {
    /* Count the number of 1 bits in the integer bits. */
    int count = 0;
    while (bits) {
        count += bits & 1;
        bits >>= 1;
    }
    return count;
}
