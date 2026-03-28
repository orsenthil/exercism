#include "luhn.h"
bool luhn(const char *num) {
    unsigned int odd_sum = 0, odd_raw_sum = 0;
    unsigned int even_sum = 0, even_raw_sum = 0;
    unsigned int zeros = 0;
    bool odd_len = false;

    while (*num != '\0') {

        if (*num >= '0' && *num <= '9')
        {
            const unsigned int n = *num - '0';
            const unsigned int doub = 2 * n;

            if (n == 0) {
                zeros += 1;
            }

            if (!odd_len)
            {
                odd_raw_sum += n;
                odd_sum += doub > 9 ? doub - 9: doub;
                odd_len = !odd_len;
            }
            else
            {
                even_raw_sum += n;
                even_sum += doub > 9 ? doub - 9: doub;
                odd_len = !odd_len;
            }
        }
        else if (*num != ' ')
            return false;

        num += 1;

    }

    const unsigned int sum = odd_len ? even_sum + odd_raw_sum : odd_sum + even_raw_sum;

    return (zeros > 1 || sum > 0) &&  sum % 10 == 0;
}
