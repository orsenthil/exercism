#include "spiral_matrix.h"

#include <stdlib.h>
#include <string.h>

spiral_matrix_t* spiral_matrix_create(int size) {
    spiral_matrix_t* sp = calloc(1, sizeof(spiral_matrix_t));
    sp->size = size;
    if (size == 0) {
        return sp;
    }
    sp->matrix = calloc(size, sizeof(int*));
    for (int i = 0; i < size; i++) {
        sp->matrix[i] = calloc(size, sizeof(int) * size);
        memset(sp->matrix[i], 0, sizeof(int) * size);
    }

    int x = 0, y = 0;
    int dx = 0, dy = 1;
    for (int i = 1; i <= size * size; i++) {
        sp->matrix[x][y] = i;
        int nx = x + dx;
        int ny = y + dy;
        if (nx < 0 || nx >= size || ny < 0 || ny >= size || sp->matrix[nx][ny] > 0) {
            int tmp = dx;
            dx = dy;
            dy = -tmp;
        }
        x += dx;
        y += dy;
    }
    return sp;
}

void spiral_matrix_destroy(spiral_matrix_t *matrix) {
    for (int i = 0; i < matrix->size; i++) {
        free(matrix->matrix[i]);
    }
    free(matrix->matrix);
    free(matrix);
}
