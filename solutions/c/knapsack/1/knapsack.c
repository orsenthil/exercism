#include "knapsack.h"
#include <stdlib.h>

#define max(x, y) (((x) > (y)) ? (x) : (y))

unsigned int maximum_value(unsigned int maximum_weight, const item_t* items, size_t item_count) {
    unsigned int* values = (unsigned int*)calloc(maximum_weight + 1, sizeof(unsigned int));
    for (size_t i = 0; i < item_count; i++) {
        for (unsigned int j = maximum_weight; j >= items[i].weight; j--) {
            values[j] = max(values[j], values[j - items[i].weight] + items[i].value);
        }
    }
    unsigned int result = values[maximum_weight];
    free(values);
    return result;
}
