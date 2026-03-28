#include "phone_number.h"

static const char DELIMITERS[] = " +().-";

#define NUM_LEN 10
#define MAX_INP_LEN 32

// legit resulting number min values 20020000
// input may have 1 as first digit but it will be discarded

char* phone_number_clean(const char* input)
{
    char* output = calloc(NUM_LEN + 1, 1);
    char phrase[MAX_INP_LEN] = {0};

    strcpy(phrase, input);

    int get_pos = 0, put_pos = 0;

    char *token = strtok(phrase, DELIMITERS);

    do {
        char c;

        for(; (c = *token); token++, get_pos++)
        {
            if (get_pos == 0 && c == '1') continue;
            if (c < '0' || c > '9') break;
            output[put_pos++] = c;
        }
    } while ((token = strtok(NULL, DELIMITERS)));

    if (put_pos < 4 || output[0]< '2' || output[3] < '2' || put_pos != 10)
    {
        memset(output, 0, put_pos);
        memset(output, '0', NUM_LEN);
    }

    output = realloc(output, NUM_LEN + 1);

    return output;
}
