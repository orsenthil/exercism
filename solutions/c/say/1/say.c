#include "say.h"
#include <stdlib.h>
#include <string.h>


static void say_below_1000(int n, const char **ones, const char **tens, char *buf) {
    if (n >= 100) {
        strcat(buf, ones[n / 100]);
        strcat(buf, " hundred");
        n %= 100;
        if (n != 0) {
            strcat(buf, " ");
        }
    }
    if ( n < 20) {
        strcat(buf, ones[n]);
    } else {
        strcat(buf, tens[n / 10]);
        n %= 10;
        if (n != 0) {
            strcat(buf, "-");
            strcat(buf, ones[n]);
        }
    }
}

int say(int64_t input, char **ans) {
    static const char *ones[] = {"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"};
    static const char *tens[] = {"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"};

    if (input < 0 || input > 999999999999) {
        return -1;
    }

    char *result = malloc(256);

    result[0] = '\0';

    if (input == 0) {
        strcat(result, "zero");
        *ans = result;
        return 0;
    }

    // loop over four groups - billions, millions, thousands and remainder
    int64_t divisors[] = {1000000000LL, 1000000LL, 1000LL, 1LL};
    const char *suffixes[] = {"billion", "million", "thousand", ""};
    for (int i = 0; i < 4; i++) {
        int group = input / divisors[i];
        if (group != 0) {
            if (result[0] != '\0') strcat(result, " ");
            say_below_1000(group, ones, tens, result);
            if (suffixes[i][0] != '\0') {
                strcat(result, " ");
                strcat(result, suffixes[i]);
            }
        }
        input %= divisors[i];
    }

    *ans = result;
    return 0;

}
