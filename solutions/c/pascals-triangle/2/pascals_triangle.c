#include "pascals_triangle.h"

void free_triangle(uint8_t** triangle, size_t rows)
{
    for (size_t i = 0; i < rows; i++) {
        free(triangle[i]);
    }
    free(triangle);
}

uint8_t** create_triangle(size_t rows)
{
    if (!rows) {
        uint8_t** triangle = calloc(1, sizeof(uint8_t* ));
        triangle[0] = calloc(1, sizeof(uint8_t));
        return triangle;
    }
    uint8_t** triangle = calloc(rows, sizeof(uint8_t*));
    triangle[0] = calloc(rows, sizeof(uint8_t));
    triangle[0][0] = 1;
    for (size_t i = 0; i < rows; i++) {
        triangle[i] = calloc(rows, sizeof(uint8_t));
        triangle[i][0] = 1;
        for (size_t j = 1; j < rows; j++) {
            triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j];
        }
    }
    return triangle;
}
