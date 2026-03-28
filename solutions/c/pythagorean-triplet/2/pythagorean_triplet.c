#include "pythagorean_triplet.h"

// check if a^2 + b ^2 = c ^ 2
// and 0 < a < b < c

bool is_pythagorean_triplet(uint16_t a, uint16_t b, uint16_t c)
{
    bool ascending = (a < b) && (b < c);
    bool pythagorean = ( a * a  + b * b == c * c);
    return ascending && pythagorean;
}

triplets_t* triplets_with_sum(uint16_t sum)
{
    triplets_t* result = malloc(sizeof(triplets_t) * sum);
    result->count = 0;

    for (uint16_t a = 1; a < sum / 3; a++)
    {
        for (uint16_t b = a + 1; b <= (sum-a)/ 2; b++)
        {
            uint16_t c = sqrt(a * a + b * b);

            if (is_pythagorean_triplet(a, b, c) && (a + b + c == sum))
            {
                triplet_t current = {.a = a, .b = b, .c = c};
                result->triplets[result->count] = current;
                result->count++;
            }
        }
    }

    return result;
}

void free_triplets(triplets_t* triplet)
{
    free(triplet);
    return;
}
