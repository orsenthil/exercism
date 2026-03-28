#include "gigasecond.h"
#include <stdint.h>
#include <math.h>

time_t gigasecond_after(time_t t)
{
    time_t result;
    result = pow(10, 9) + t;
    return result;
}
