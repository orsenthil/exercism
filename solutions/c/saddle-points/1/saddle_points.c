#include "saddle_points.h"
#include <stdlib.h>
#include <stdbool.h>

static bool is_saddle_point(size_t rows, size_t cols, uint8_t matrix[rows][cols], size_t row, size_t col) {
    uint8_t value = matrix[row][col];
    bool is_saddle_point = true;
    for (size_t i = 0; i < cols; i++) {
        if (matrix[row][i] > value) {
            is_saddle_point = false;
            break;
        }
    }
    for (size_t i = 0; i < rows; i++) {
        if (matrix[i][col] < value) {
            is_saddle_point = false;
            break;
        }
    }
    return is_saddle_point;
}

saddle_points_t *saddle_points(size_t rows, size_t cols, uint8_t matrix[rows][cols]) {
    saddle_points_t *saddle_points = malloc(sizeof(saddle_points_t));
    saddle_points->count = 0;
    saddle_points->points = NULL;
    for (size_t row = 0; row < rows; row++) {
        for (size_t col = 0; col < cols; col++) {
            if (is_saddle_point(rows, cols, matrix, row, col)) {
                saddle_points->count++;
                saddle_points->points = realloc(saddle_points->points, saddle_points->count * sizeof(saddle_point_t));
                saddle_points->points[saddle_points->count - 1].row = row + 1;
                saddle_points->points[saddle_points->count - 1].column = col + 1;
            }
        }
    }
    return saddle_points;
}
void free_saddle_points(saddle_points_t *saddle_points) {
    free(saddle_points->points);
    free(saddle_points);
}

