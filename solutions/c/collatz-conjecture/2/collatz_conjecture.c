#include "collatz_conjecture.h"

int steps(int start)
{
    int count = 0;
    if (start == 1) {
        return count;
    }

    while (start > 1)
    {
        if (start % 2 == 0) {
            start = start / 2;
        } else {
            start = (start * 3) + 1;
        }
        if (start == 1) {
            count += 1;
            return count;
        } else {
            count += 1;
        }
    }
    return -1;
}
