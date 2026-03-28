#include "phone_number.h"

char* phone_number_clean(const char* input) {
    char* output;
    output = malloc(sizeof (input));
    strcpy(output, input);
    return output;
}
