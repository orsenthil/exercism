#include "allergies.h"
#include <string.h>
#include <stdio.h>

/*
 * Subtract numbers until you exhaust the score
 * flag the subtracted number along with increasing count.
 */

static int find_less_or_equal(int *arr, int len, int num)
{
    // find a index in the array which is less than or equal to number

    for (int i = len-1; i >= 0; i--)
    {
       if (arr[i] <= num)
           return i;
    }
    return -1;
}

allergen_list_t get_allergens(uint16_t score)
{
    int scores[] = {1, 2, 4, 8, 16, 32, 64, 128};
    allergen_list_t my_list;
    memset(my_list.allergens, 0, sizeof(my_list.allergens));
    my_list.count = 0;
    int prev_idx = -1;
    int tries = 0;

    // skipping 256, 512, 1024

    if (score > 128)
    {
        for (int i = 2; i < 10; i++)
            if (score == 128 * i)
                return my_list;
    }

    // cut down the upper part of the score

    int temp = score % 256;
    printf("%d\n", temp);

    while (tries < 100 || temp < 0)
    {
        int idx = find_less_or_equal(scores, ALLERGEN_COUNT, temp);

        if (idx != -1 && prev_idx != idx)
        {
            if (!my_list.allergens[idx]) {
                my_list.allergens[idx] = 1;
                my_list.count++;
            }

            temp -= scores[idx];
        }

        prev_idx = idx;
        tries++;
    }

    if (temp) {
        // returning eggs for any unfinished value
        memset(my_list.allergens, 0, sizeof(my_list.allergens));
        my_list.count = 1;
        my_list.allergens[0] = 1;
    }

    return my_list;
}

bool is_allergic_to(allergen_t allergen, uint16_t score) {
    allergen_list_t  result = get_allergens(score);

    if (result.allergens[allergen]) {
        return true;
    }

    return false;

}
