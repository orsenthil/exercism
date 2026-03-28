#include "sublist.h"

#include <stdbool.h>
#include <stdlib.h>
#include <string.h>

static bool issublist(int *smaller, int *larger, size_t smaller_length, size_t larger_length)
{
    if (smaller_length == 0) {
        return true;
    }
    for (size_t i = 0; i <= larger_length - smaller_length; i++) {
        if (memcmp(smaller, larger + i, sizeof(int) * smaller_length) == 0) {
            return true;
        }
    }
    return false;
}

comparison_result_t check_lists(int *list1, int *list2,
                                size_t list1_length,
                                size_t list2_length) {

    if (list1_length == list2_length && issublist(list1, list2, list1_length, list2_length)) {
        return EQUAL;
    }

    if (list1_length < list2_length) {
        return issublist(list1, list2, list1_length, list2_length) ? SUBLIST: UNEQUAL;
    }

    return issublist(list2, list1, list2_length, list1_length) ? SUPERLIST: UNEQUAL;

}
