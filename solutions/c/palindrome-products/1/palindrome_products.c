#include "palindrome_products.h"
#include <stdio.h>
#include <stdlib.h>
#include <stddef.h>
#include <stdbool.h>
#include <limits.h>

static int is_palindrome(int n) {
    int original = n;
    int reversed = 0;
    while (n > 0) {
        reversed = reversed * 10 + n % 10;
        n /= 10;
    }
    return original == reversed;
}

void free_product(product_t *p) {
    factor_t *current = p->factors_sm;
    while (current != NULL) {
        factor_t *next = current->next;
        free(current);
        current = next;
    }
    current = p->factors_lg;
    while (current != NULL) {
        factor_t *next = current->next;
        free(current);
        current = next;
    }
    free(p);
}

product_t *get_palindrome_product(int from, int to) {
    product_t *product = malloc(sizeof(product_t));
    product->smallest = INT_MAX;
    product->largest = INT_MIN;
    product->error[0] = '\0';
    product->factors_sm = NULL;
    product->factors_lg = NULL;

    if (from > to) {
        snprintf(product->error, MAXERR, "invalid input: min is %d and max is %d", from, to);
        return product;
    }

    for (int i = from; i <= to; i++) {
        for (int j = i; j <= to; j++) {
            int value = i * j;
            if (is_palindrome(value)) {
                if (value < product->smallest) {
                    // free the current factors_sm list
                    factor_t *current = product->factors_sm;
                    while (current != NULL) {
                        factor_t *next = current->next;
                        free(current);
                        current = next;
                    }
                    product->smallest = value;
                    product->factors_sm = malloc(sizeof(factor_t));
                    product->factors_sm->factor_a = i;
                    product->factors_sm->factor_b = j;
                    product->factors_sm->next = NULL;
                } else if (value == product->smallest) {
                    factor_t *node = malloc(sizeof(factor_t));
                    node->factor_a = i;
                    node->factor_b = j;
                    node->next = product->factors_sm;
                    product->factors_sm = node;
                } 
                
                if (value > product->largest) {
                    // free the current factors_lg list
                    factor_t *current = product->factors_lg;
                    while (current != NULL) {
                        factor_t *next = current->next;
                        free(current);
                        current = next;
                    }
                    product->largest = value;
                    product->factors_lg = malloc(sizeof(factor_t));
                    product->factors_lg->factor_a = i;
                    product->factors_lg->factor_b = j;
                    product->factors_lg->next = NULL;
                } else if (value == product->largest) {
                    factor_t *node = malloc(sizeof(factor_t));
                    node->factor_a = i;
                    node->factor_b = j;
                    node->next = product->factors_lg;
                    product->factors_lg = node;
                }
            }
        }
    }

    if (product->smallest == INT_MAX) {
        snprintf(product->error, MAXERR, "no palindrome with factors in the range %d to %d", from, to);
        return product;
    }

    if (product->largest == INT_MIN) {
        snprintf(product->error, MAXERR, "no palindrome with factors in the range %d to %d", from, to);
        return product;
    }

    return product;
}
