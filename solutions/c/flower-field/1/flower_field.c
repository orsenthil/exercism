#include "flower_field.h"
#include <stdlib.h>
#include <string.h>

char **annotate(const char **garden, const size_t rows) {
    if (rows == 0) {
        return NULL;
    }
    char **annotation = malloc((rows + 1) * sizeof(char *));
    for (size_t i = 0; i < rows; i++) {
        annotation[i] = malloc(strlen(garden[i]) + 1);
        strcpy(annotation[i], garden[i]);
    }
    for (size_t i = 0; i < rows; i++) {
        size_t cols = strlen(garden[i]);
        for (size_t j = 0; j < cols; j++) {
            if (garden[i][j] == '*') {
                annotation[i][j] = '*';
            } else {
                int count = 0;
                for (int di = -1; di <= 1; di++) {
                    for (int dj = -1; dj <= 1; dj++) {
                        if (di == 0 && dj == 0) {
                            continue;
                        }
                        int ni = (int)i + di;
                        int nj = (int)j + dj;
                        if (ni >= 0 && ni < (int)rows && nj >= 0 && nj < (int)cols) {
                            if (garden[ni][nj] == '*') {
                                count++;
                            }
                        }
                    }
                }
                if (count > 0) {
                    annotation[i][j] = '0' + count;
                }
            }
        }
    }

    annotation[rows] = NULL;
    return annotation;
}
void free_annotation(char **annotation) {
    for (size_t i = 0; annotation[i] != NULL; i++) {
        free(annotation[i]);
    }
    free(annotation);
}
