#include "matching_brackets.h"
#include <string.h>
#include <stdlib.h>

bool is_paired(const char *input) {
    char stack[strlen(input)];
    size_t top = 0;

    for (size_t i = 0; i < strlen(input); i++) {
        char c = input[i];
        if (c == '(' || c == '[' || c == '{') {
            stack[top++] = c;
        } else if (c == ')' || c == ']' || c == '}') {
            if (top == 0) {
                return false;
            }
            char last = stack[top - 1];
            if (last == '(' && c == ')') {
                top--;
            } else if (last == '[' && c == ']') {
                top--;
            } else if (last == '{' && c == '}') {
                top--;
            } else {
                return false;
            }
        }
    }
    return top == 0;
}
