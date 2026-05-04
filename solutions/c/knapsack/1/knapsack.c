#include "knapsack.h"
#include <stdlib.h>

static unsigned int max(unsigned int a, unsigned int b) {
    return a > b ? a : b;
}

unsigned int maximum_value(unsigned int maximum_weight, item_t *items, size_t item_count) {
    if (item_count == 0 || maximum_weight == 0) {
        return 0;
    }

    // create a 2D array to store the maximum value for each weight and item
    int dp[item_count + 1][maximum_weight + 1];
    for (size_t i = 0; i <= item_count; i++) {
        for (size_t j = 0; j <= maximum_weight; j++) {
            dp[i][j] = 0;
        }
    }

    // fill the dp array
    for (size_t i = 1; i <= item_count; i++) {
        for (size_t j = 1; j <= maximum_weight; j++) {
            if (items[i - 1].weight <= j) {
                dp[i][j] = max(dp[i - 1][j], dp[i - 1][j - items[i - 1].weight] + items[i - 1].value);
            } else {
                dp[i][j] = dp[i - 1][j];
            }
        }
    }
    return dp[item_count][maximum_weight];
}
