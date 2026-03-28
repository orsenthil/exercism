#include "gigasecond.h"
#include <stdint.h>
#include <math.h>

time_t gigasecond_after(time_t t)
{
    time_t result;
    uint64_t giga = pow(10, 9) + t;
    result = giga % 60;
    return result;
}
