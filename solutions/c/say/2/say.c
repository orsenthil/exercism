#include "say.h"


int say(int64_t input, char **ans)
{
    if (input > 999999999999 || input < 0)
        return -1;

    char buffer[BUFFER_SIZE] = { 0 };

    if (!input)
        strcpy(buffer, "zero");

    char* one_to_ten[] = {"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"};
    char* ten_to_twenty[] = {"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"};
    char* tens[] = {"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"};
    char* scales[] = {"thousand", "million", "billion"};

    for (int8_t i = (int8_t) log10(input) / 3 * 3; i > -1 && input; i -= 3) {
        uint64_t pow_res = pow(10, i);
        uint16_t val = input / pow_res;
        input %= pow_res;

        while (val) {
            if (val > 99) {
                strcat(buffer, one_to_ten[(uint8_t)(val / 100) - 1]);
                strcat(buffer, " hundred");
                val %= 100;
                if (val)
                    strcat(buffer, " ");
            }
            else if (val > 19) {
                strcat(buffer, tens[val / 10 - 2]);
                val %= 10;
                if (val)
                    strcat(buffer, "-");
            }
            else if (val > 9) {
                strcat(buffer, ten_to_twenty[val % 10]);
                val = 0;
            }
            else {
                strcat(buffer, one_to_ten[val - 1]);
                val = 0;
            }
        }

        if (i) {
            strcat(buffer, " ");
            strcat(buffer, scales[i / 3 - 1]);
            if (input)
                strcat(buffer, " ");
        }
    }

    *ans = calloc(BUFFER_SIZE, 1);
    strcpy(*ans, buffer);

    return 0;
}
