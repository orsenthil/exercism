#include "spiral_matrix.h"
#include <stdlib.h>
#include <stddef.h>

spiral_matrix_t *spiral_matrix_create(int size) {
    if (size == 0) {
        spiral_matrix_t *spiral_matrix = malloc(sizeof(spiral_matrix_t));
        spiral_matrix->size = 0;
        spiral_matrix->matrix = NULL;
        return spiral_matrix;
    }
    spiral_matrix_t *spiral_matrix = malloc(sizeof(spiral_matrix_t));
    spiral_matrix->size = size;
    spiral_matrix->matrix = malloc(sizeof(int *) * size);
    for (int i = 0; i < size; i++) {
        spiral_matrix->matrix[i] = malloc(sizeof(int) * size);
    }
    int num = 1;
    int top = 0, bottom = size - 1, left = 0, right = size - 1;
    while (top <= bottom && left <= right) {
        for (int i = left; i <= right; i++) {
            spiral_matrix->matrix[top][i] = num++;
        }
        top++;
        for (int i = top; i <= bottom; i++) {
            spiral_matrix->matrix[i][right] = num++;
        }
        right--;
        for (int i = right; i >= left; i--) {
            spiral_matrix->matrix[bottom][i] = num++;
        }
        bottom--;
        for (int i = bottom; i >= top; i--) {
            spiral_matrix->matrix[i][left] = num++;
        }
        left++;
    }
    return spiral_matrix;
}

void spiral_matrix_destroy(spiral_matrix_t *spiral_matrix) {
    if (spiral_matrix == NULL) {    
        return;
    }
    for (int i = 0; i < spiral_matrix->size; i++) {
        free(spiral_matrix->matrix[i]);
    }
    free(spiral_matrix->matrix);
    free(spiral_matrix);
}
