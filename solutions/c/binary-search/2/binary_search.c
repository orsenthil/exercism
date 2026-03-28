#include "binary_search.h"

int *binary_search(int value, const int *arr, size_t length) {
    int left = 0;
    int right = length - 1;

    if (length == 0 || arr == NULL) {
        return NULL;
    }

    while (left <= right) {
        int mid = left + (right - left) / 2;

        if (value == arr[mid])
            return (int*)&arr[mid];

        if (value > arr[mid]) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    return NULL;
}
