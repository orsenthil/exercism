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
    if (rows == 0) {
        uint8_t** triangle = malloc(sizeof(uint8_t*));
        triangle[0] = malloc(sizeof(uint8_t));
        triangle[0][0] = 0;
        return triangle;
    }

    uint8_t **triangle = malloc(rows * sizeof(uint8_t*));

    for (size_t i = 0; i < rows; i++) {
        triangle[i] = malloc((rows + 1) * sizeof(uint8_t));
    }

    for (size_t i = 0; i < rows; i++) {
        for (size_t j = 0; j < rows; j++) {
            triangle[i][j] = 0;
        }
    }

    for(size_t i=0; i < rows; i++) {
        if (i == 0) {
            triangle[i][0] = 1;
        }

        if (i > 0) {
            for (size_t j = 0; j <= i; j++)
            {
                if (j == 0) {
                    triangle[i][j] = triangle[i-1][j];
                }
                else {
                    triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j];
                }
            }
        }
    }
    return triangle;
}
