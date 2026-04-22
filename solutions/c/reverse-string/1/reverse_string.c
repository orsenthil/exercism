#include "reverse_string.h"
#include <string.h>
#include <stdlib.h>


char *reverse(const char *value) {
    if (value == NULL) {
        return NULL;
    }
    size_t len = strlen(value);
    char * buf =malloc(len + 1);
    if (buf == NULL) {
        return NULL;
    }
    for (size_t i = 0; i < len; i++) {
        buf[i] = value[len - 1 - i];
    }
    buf[len] = '\0';
    return buf;
}
