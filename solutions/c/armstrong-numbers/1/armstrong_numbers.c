#include "armstrong_numbers.h"
#include <math.h>

bool is_armstrong_number(int candidate) {
  if (candidate == 0) {
    return true;
  }
  int number = candidate;
  int digits = 0;
  int result = 0;

  while (number > 0) {
    digits++;
    number /= 10;
  }

  number = candidate;
  while (number > 0) {
    int last_digit = number % 10;
    result += pow(last_digit, digits);

    number /= 10;
  }

  return result == candidate;
}
