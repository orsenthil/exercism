#include "saddle_points.h"

#include <stdio.h>

saddle_points_t* saddle_points(size_t nrows, size_t ncols, uint8_t matrix[nrows][ncols]) {
    int row_max[nrows];
    int col_min[ncols];
    int val;

    for(size_t i = 0; i < nrows; i++) {
        for(size_t j = 0; j < ncols; j++) {
            val = matrix[i][j];

            if (j == 0) {
                row_max[i] = val;
            }

            if (i == 0) {
                col_min[j] = val;
            }

            if (val >= row_max[i]) {
                row_max[i] = val;
            }

            if (val <= col_min[j]) {
                col_min[j] = val;
            }
        }
    }

    // Find Saddle Points

    saddle_point_t  *points = calloc(1, sizeof(saddle_point_t));
    size_t count = 0;

    for (size_t i = 0; i < nrows; i++)
    {
        for (size_t j = 0; j < ncols; j++)
        {
            val = matrix[i][j];

            if (val >= row_max[i] && val <= col_min[j])
            {
                points[count].row = i + 1;
                points[count].column = j + 1;
                count++;

                points = realloc(points, sizeof(saddle_point_t) * (count + 1));
            }
        }
    }

    saddle_points_t *saddle = calloc(1, sizeof(saddle_point_t) + sizeof(saddle_point_t) * count);

    saddle->count = count;

    for (size_t i = 0; i < count; i++)
    {
        saddle->points[i].row = points[i].row;
        saddle->points[i].column = points[i].column;
    }

    return saddle;
}


void free_saddle_points(saddle_points_t *saddle) {
    free(saddle);
}
