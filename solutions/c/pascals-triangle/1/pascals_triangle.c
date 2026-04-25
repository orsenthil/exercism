#include "pascals_triangle.h"
#include <stdlib.h>

uint8_t **create_triangle(size_t rows) {
    uint8_t **triangle = malloc((rows ? rows : 1) * sizeof(uint8_t *));

    if (rows == 0) {
        triangle[0] = calloc(1, sizeof(uint8_t));
        return triangle;
    }

    for (size_t i = 0; i < rows; i++) {
        triangle[i] = calloc(rows, sizeof(uint8_t));
        triangle[i][0] = 1;
        triangle[i][i] = 1;
    }

    for (size_t i = 2; i < rows; i++) {
        for (size_t j = 1; j < i; j++) {
            triangle[i][j] = triangle[i - 1][j - 1] + triangle[i - 1][j];
        }
    }
    return triangle;
}


void free_triangle(uint8_t **triangle, size_t rows) {
    for (size_t i = 0; i < rows; i++) {
        free(triangle[i]);
    }
    free(triangle);
}
