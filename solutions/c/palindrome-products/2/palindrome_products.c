#include "palindrome_products.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static bool is_palindrome(int number) {
    int reverse = 0, value = number;
    while (value != 0) {
        reverse *= 10;
        reverse += value % 10;
        value /= 10;
    }
    return reverse == number;
}

static void clear(factor_t *factors) {
    if (factors == NULL) {
        return;
    }
    factor_t *current = factors->next;
    while (current != NULL) {
        factor_t *next = current->next;
        free(current);
        current = next;
    }
}

static void set(factor_t *factors, int factor_a, int factor_b) {
    clear(factors);
    factors->factor_a = factor_a;
    factors->factor_b = factor_b;
    factors->next = NULL;
}

static void add(factor_t *factors, int factor_a, int factor_b) {
    factor_t *current = factors;
    while (current->next) {
        current = current->next;
    }
    factor_t *new_factors = calloc(1, sizeof(factor_t));
    new_factors->factor_a = factor_a;
    new_factors->factor_b = factor_b;
    new_factors->next = NULL;
    current->next = new_factors;
}


product_t *get_palindrome_product(int from, int to) {
    product_t *result = calloc(1, sizeof(product_t));

    if (to < from) {
        sprintf(result->error, "invalid input: min is %d and max is %d", from, to);
        return result;
    }

    factor_t *min_factors = calloc(1, sizeof(factor_t));
    factor_t *max_factors = calloc(1, sizeof(factor_t));

    int min_product = to * to;
    int max_product = 0;

    for (int i = from; i <= to; i++) {
        for (int j = i; j <= to; j++) {
            int product = i * j;
            if (is_palindrome(product)) {
                if (product < min_product) {
                    set(min_factors, i, j);
                    min_product = product;
                } else if (product == min_product) {
                    add(min_factors, i, j);
                }
                if (product > max_product) {
                    set(max_factors, i, j);
                    max_product = product;
                } else if (product == max_product) {
                    add(max_factors, i, j);
                }
            }
        }
    }

    if (min_factors->factor_a == 0) {
        sprintf(result->error, "no palindrome with factors in the range %d to %d", from, to);
    } else {
        result->factors_sm = min_factors;
        result->factors_lg = max_factors;
        result->smallest = min_product;
        result->largest = max_product;
    }
    return result;
}

void free_product(product_t *p) {
    clear(p->factors_sm);
    clear(p->factors_lg);
    free(p->factors_sm);
    free(p->factors_lg);
    free(p);
}
