#include "matching_brackets.h"
#include <stdlib.h>

bool is_paired(const char *input)
{
    char expected[50] = "";
    size_t size = 0;
    for(const char *chr = input; *chr; chr++)
    {
        switch (*chr) {
            case '(':
                expected[size] = ')';
                size += 1;
                break;
            case '{':
                expected[size] = '}';
                size += 1;
                break;
            case '[':
                expected[size] = ']';
                size += 1;
                break;
            case ')':
            case '}':
            case ']':
                size -= 1;
                if (expected[size] != *chr) {
                    return false;
                }
                expected[size] = 0;
                break;
            default:
                break;

        }
    }
    return size == 0;
}
