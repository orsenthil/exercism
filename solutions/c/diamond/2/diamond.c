#include "diamond.h"
#include <stdlib.h>
#include <string.h>

char **make_diamond(const char letter) {
    int n = letter - 'A' + 1;
    int width = 2 * n - 1;
    int rows = 2 * n - 1;
    char **diamond = malloc(rows * sizeof(char *));
    for (int i = 0; i < n; i++) {
        diamond[i] = malloc(width * sizeof(char));
        for (int j = 0; j < width; j++) {
            diamond[i][j] = ' ';
        }
        diamond[i][n -1 - i] = 'A' + i;
        if (i > 0) {
            diamond[i][n - 1+ i] = 'A' + i;
        }
    }
    for (int i = 0; i < n - 1; i++) {
        diamond[i][width] = '\0';
        diamond[rows - 1 - i] = malloc(width + 1);
        memcpy(diamond[rows - 1 - i], diamond[i], width + 1);
    }
    return diamond;
}

void free_diamond(char **diamond) {
    int rows = strlen(diamond[0]);
    for (int i = 0; i < rows; i++) {
        free(diamond[i]);
    }
    free(diamond);
}
