#ifndef PANGRAM_H
#define PANGRAM_H

#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

bool is_pangram(const char *sentence);

bool all(const int* alphabet, size_t n);
char* to_lower(const char* sentence);

#endif
