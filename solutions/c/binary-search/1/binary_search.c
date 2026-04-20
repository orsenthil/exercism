#include "binary_search.h"

const int *binary_search(int value, const int *arr, size_t length) {
  if (length == 0) {
    return NULL;
  }
  int left = 0;
  int right = length - 1;
  int mid = 0;

  while (left <= right) {
    mid = left + (right - left) / 2;

    if (arr[mid] == value) {
      return &arr[mid];
    }
    if (arr[mid] < value) {
      left = mid + 1;
    }
    if (arr[mid] > value) {
      right = mid - 1;
    }
  }
  return NULL;
}
