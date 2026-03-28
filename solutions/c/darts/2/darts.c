#include "darts.h"
#include <math.h>

int score(coordinate_t position) {
    float distance = sqrt(pow(position.x, 2) + pow(position.y, 2));

    if (distance > 10)
    {
        return 0;
    }
    else if (distance > 5)
    {
        return 1;
    }
    else if (distance > 1 )
    {
        return 5;
    }

    return 10;
}