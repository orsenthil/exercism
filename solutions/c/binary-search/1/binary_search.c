#include "binary_search.h"

int *binary_search(int value, const int *arr, size_t length) {
    bool found = false;
    size_t left = 0;
    size_t right = length-1;
    size_t mid = 0;

    while (found == false && left < right) {
        mid = left + ((right - left) / 2);
        if (arr[mid] == value) {
            found = true;
        } else if (value < arr[mid]) {
            right = mid - 1;
        } else if (value > arr[mid]) {
            left = mid + 1;
        }
    }
    if (found == false) {
        return NULL;
    }


    return arr[mid];
}