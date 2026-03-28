#include "minesweeper.h"

#include <stdlib.h>
#include <string.h>


static size_t last_rows = 0;

char **annotate(const char **minefield, const size_t rows)
{
    if (minefield == NULL)
        return NULL;

    last_rows = rows;

    size_t row_length = strlen(minefield[0]);
    char** result = calloc(rows, sizeof(char *));
    for(size_t i = 0; i < rows; i++)
    {
        result[i] = calloc(row_length + 1, sizeof(char));
        for(size_t j = 0; j < row_length; j++) {
            result[i][j] = '0';
        }
        result[i][row_length] = '\0';
    }

    for (size_t i = 0; i < rows; i++) {
        for(size_t j = 0; j < row_length; j++) {
            if (minefield[i][j] == '*') {
                result[i][j] = '*';
                // One Row up
                if (i > 0) {
                    if (j > 0 && minefield[i - 1][j - 1] != '*')
                        result[i - 1][j - 1]++;
                    //Center
                    if (minefield[i - 1][j] != '*')
                        result[i - 1][j]++;
                    // Right
                    if (j + 1 < row_length && minefield[i - 1][j + 1] != '*')
                        result[i - 1][j + 1]++;
                }
                // Left
                if (j > 0 && minefield[i][j - 1] != '*')
                    result[i][j - 1]++;
                // Right
                if (j + 1 < row_length && minefield[i][j + 1] != '*')
                    result[i][j + 1]++;

                // One Row Down

                if (i + 1 < rows) {
                    if (j > 0 && minefield[i + 1][j - 1] != '*')
                        result[i + 1][j - 1]++;

                    if (minefield[i + 1][j] != '*')
                        result[i + 1][j]++;

                    if (j + 1 < row_length && minefield[i + 1][j + 1] != '*')
                        result[i + 1][j + 1]++;
                }
            }
        }
    }

    // Replace 0 mines with whitespace

    for(size_t i = 0; i < rows; i++)
        for(size_t j = 0; j < row_length; j++)
            if (result[i][j] == '0')
                result[i][j] = ' ';

    return result;
}
